package clickhouse

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2/lib/driver"
	"google.golang.org/protobuf/types/known/structpb"

	ubaV1 "go-wind-uba/api/gen/go/uba/service/v1"
)

const (
	// adhocQueryTimeout 单条自助查询的服务端超时。
	adhocQueryTimeout = 30 * time.Second

	// adhocQueryMaxRows 行数上限硬顶。
	adhocQueryMaxRows = 1000
)

// readOnlyKeywordWhitelist 允许的语句首关键字（只读白名单）。
var readOnlyKeywordWhitelist = []string{
	"SELECT", "SHOW", "DESC", "DESCRIBE", "EXPLAIN",
}

// forbiddenSQLMarkers 自助查询里不允许出现的片段：多语句、注释都可能改变语句语义。
var forbiddenSQLMarkers = []string{";", "--", "/*", "*/", "#"}

// ValidateAdhocQuery 只读查询静态校验：非空、单语句、无注释、首关键字在白名单。
func ValidateAdhocQuery(query string) error {
	trimmed := strings.TrimSpace(query)
	if trimmed == "" {
		return ubaV1.ErrorBadRequest("query is empty")
	}

	// 容忍结尾的单个分号，其余分号（多语句）一律拒绝
	trimmed = strings.TrimSuffix(trimmed, ";")
	if strings.Contains(trimmed, ";") {
		return ubaV1.ErrorBadRequest("multiple statements are not allowed")
	}

	lower := strings.ToLower(trimmed)
	for _, marker := range forbiddenSQLMarkers {
		if strings.Contains(lower, marker) {
			return ubaV1.ErrorBadRequest("comment markers are not allowed: %q", marker)
		}
	}

	first := strings.ToUpper(strings.Fields(trimmed)[0])
	allowed := false
	for _, kw := range readOnlyKeywordWhitelist {
		if first == kw {
			allowed = true
			break
		}
	}
	if !allowed {
		return ubaV1.ErrorBadRequest("only read-only queries (SELECT/SHOW/DESCRIBE/EXPLAIN) are allowed, got: %s", first)
	}

	return nil
}

// AdhocRowQuerier 自助查询需要的原始行集通道。crud 封装的 Select 只能扫进结构体，
// 而自助查询的列在执行前未知，必须走 clickhouse-go 的原始行集按列名动态取值。
// clickhouseV2.Conn 满足该接口。
type AdhocRowQuerier interface {
	Query(ctx context.Context, query string, args ...any) (driver.Rows, error)
}

// ExecuteAdhocQuery 执行自助只读查询：带超时与行数上限，结果转 JSON 安全值。
func (r *AnalyticsRepo) ExecuteAdhocQuery(
	ctx context.Context,
	query string,
	maxRows int32,
) (*ubaV1.ExecuteSqlQueryResponse, error) {
	if err := ValidateAdhocQuery(query); err != nil {
		return nil, err
	}
	if r.adhoc == nil {
		return nil, ubaV1.ErrorInternalServerError("clickhouse ad-hoc connection is not available")
	}
	if maxRows <= 0 {
		maxRows = 200
	}
	if maxRows > adhocQueryMaxRows {
		maxRows = adhocQueryMaxRows
	}

	runCtx, cancel := context.WithTimeout(ctx, adhocQueryTimeout)
	defer cancel()

	start := time.Now()

	rows, err := r.adhoc.Query(runCtx, strings.TrimSpace(query))
	if err != nil {
		r.log.Errorf("adhoc query failed: %v", err)
		return nil, ubaV1.ErrorInternalServerError("query failed: %v", err)
	}
	defer rows.Close()

	columns := rows.Columns()

	// 多读一行用于判定截断
	resultRows := make([]*ubaV1.SqlResultRow, 0, maxRows)
	truncated := false
	for rows.Next() {
		if int32(len(resultRows)) >= maxRows {
			truncated = true
			break
		}

		values := make([]any, len(columns))
		scan := make([]any, len(columns))
		for i := range values {
			scan[i] = &values[i]
		}
		if err := rows.Scan(scan...); err != nil {
			r.log.Errorf("adhoc query scan failed: %v", err)
			return nil, ubaV1.ErrorInternalServerError("query scan failed")
		}

		cells := make([]*structpb.Value, 0, len(columns))
		for _, cell := range values {
			v, convErr := chToStructValue(cell)
			if convErr != nil {
				return nil, ubaV1.ErrorInternalServerError("query result conversion failed")
			}
			cells = append(cells, v)
		}
		resultRows = append(resultRows, &ubaV1.SqlResultRow{Values: cells})
	}
	if err := rows.Err(); err != nil {
		r.log.Errorf("adhoc query iterate failed: %v", err)
		return nil, ubaV1.ErrorInternalServerError("query failed")
	}

	return &ubaV1.ExecuteSqlQueryResponse{
		Columns:    columns,
		Rows:       resultRows,
		RowCount:   uint64(len(resultRows)),
		Truncated:  truncated,
		DurationMs: time.Since(start).Milliseconds(),
	}, nil
}

// chToStructValue 把驱动返回的任意标量/复合值转为 JSON 安全的 protobuf Value。
// 驱动专有类型分四层兜底：nil/时间/字节串 → 切片/Map 反射 → Stringer → fmt.Sprint。
func chToStructValue(v any) (*structpb.Value, error) {
	switch val := v.(type) {
	case nil:
		return structpb.NewNullValue(), nil
	case time.Time:
		return structpb.NewStringValue(val.Format(time.RFC3339Nano)), nil
	case []byte:
		return structpb.NewStringValue(string(val)), nil
	case string, bool,
		int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64,
		float32, float64:
		return structpb.NewValue(val)
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		items := make([]*structpb.Value, rv.Len())
		for i := range items {
			item, err := chToStructValue(rv.Index(i).Interface())
			if err != nil {
				return nil, err
			}
			items[i] = item
		}
		return structpb.NewListValue(&structpb.ListValue{Values: items}), nil
	case reflect.Map:
		fields := make(map[string]*structpb.Value, rv.Len())
		iter := rv.MapRange()
		for iter.Next() {
			key := iter.Key().Interface()
			ks, ok := key.(string)
			if !ok {
				ks = fmt.Sprint(key)
			}
			item, err := chToStructValue(iter.Value().Interface())
			if err != nil {
				return nil, err
			}
			fields[ks] = item
		}
		return structpb.NewStructValue(&structpb.Struct{Fields: fields}), nil
	}

	// Decimal/UUID/IP 等驱动专有类型普遍实现 Stringer，降级为字符串展示。
	if s, ok := v.(fmt.Stringer); ok {
		return structpb.NewStringValue(s.String()), nil
	}
	return structpb.NewStringValue(fmt.Sprint(v)), nil
}
