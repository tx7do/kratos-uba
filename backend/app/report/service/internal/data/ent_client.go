package data

import (
	"context"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/go-utils/entgo"
	"github.com/tx7do/kratos-bootstrap/gen/api/go/conf/v1"

	"kratos-uba/app/report/service/internal/data/ent"
	"kratos-uba/app/report/service/internal/data/ent/migrate"
)

// NewEntClient 创建Ent ORM数据库客户端
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *entgo.EntClient[*ent.Client] {
	l := log.NewHelper(log.With(logger, "module", "ent/data/report-service"))

	drv, err := entgo.CreateDriver(cfg.Data.Database.Driver, cfg.Data.Database.Source,
		int(cfg.Data.Database.MaxIdleConnections), int(cfg.Data.Database.MaxOpenConnections), cfg.Data.Database.ConnectionMaxLifetime.AsDuration())
	if err != nil {
		l.Fatal(err.Error())
		return nil
	}

	client := ent.NewClient(ent.Driver(drv))
	if client == nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}

	// 运行数据库迁移工具
	if cfg.Data.Database.Migrate {
		if err = client.Schema.Create(context.Background(), migrate.WithForeignKeys(true)); err != nil {
			l.Fatalf("failed creating schema resources: %v", err)
		}
	}

	return entgo.NewEntClient(client, drv)
}
