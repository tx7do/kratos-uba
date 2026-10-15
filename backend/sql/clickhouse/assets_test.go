package clickhouse

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"testing"
	"unicode/utf8"

	"go-wind-uba/pkg/dorisinit"
)

// The scripts in this directory are the input to pkg/clickhouseinit. These tests are the
// guard rail that keeps them parseable: a byte-order mark before the leading "--" makes
// ClickHouse reject the whole file, a CRLF leaks a stray "\r" into every statement, and an
// unrendered placeholder reaches the server as a syntax error that hides its own cause.
// (The doris package holds the same guard rail for its side of sql/.)

var legacyPlaceholder = regexp.MustCompile(`\$\{[A-Za-z_][A-Za-z0-9_]*\}`)

func TestShippedSQLIsCleanUTF8(t *testing.T) {
	var checked int
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() || !strings.HasSuffix(path, ".sql") {
			return err
		}
		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		checked++

		if bytes.HasPrefix(src, []byte{0xEF, 0xBB, 0xBF}) {
			t.Errorf("%s: starts with a UTF-8 BOM, which breaks the first '--' comment", path)
		}
		if i := bytes.IndexByte(src, '\r'); i >= 0 {
			line := 1 + bytes.Count(src[:i], []byte{'\n'})
			t.Errorf("%s:%d: CR found; SQL scripts must be LF (see the root .gitattributes)", path, line)
		}
		if m := legacyPlaceholder.Find(src); m != nil {
			t.Errorf("%s: %s is not a supported placeholder; use {{.Name}} so uba-ingest can render it", path, m)
		}
		if !utf8.Valid(src) {
			t.Errorf("%s: not valid UTF-8", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if checked < 4 {
		t.Fatalf("only %d .sql files checked from %s; the walk is broken", checked, ".")
	}
}

func TestInitScriptsRenderAndSplit(t *testing.T) {
	params := dorisinit.Params{"KafkaBrokerList": "kafka:9092"}

	var empty []string
	brokerRendered := false

	for _, s := range InitScripts() {
		src, err := dorisinit.Render(s.Name, s.Bytes, params)
		if err != nil {
			t.Fatalf("%s: %v", s.Name, err)
		}
		stmts, err := dorisinit.SplitStatements(src)
		if err != nil {
			t.Fatalf("%s: %v", s.Name, err)
		}
		if len(stmts) == 0 {
			empty = append(empty, s.Name)
		}
		for _, st := range stmts {
			if strings.Contains(st.Text, "{{") {
				t.Errorf("%s:%d has an unrendered placeholder:\n%s", s.Name, st.Line, st.Text)
			}
			if strings.Contains(st.Text, "kafka_broker_list = 'kafka:9092'") {
				brokerRendered = true
			}
		}
	}

	// 04_indexes.sql is comments-only today. It stays in the apply list on purpose, but if
	// it ever gains statements this test should be told about it rather than go stale.
	if want := []string{"04_indexes.sql"}; !slices.Equal(empty, want) {
		t.Errorf("scripts with no statements = %v, want %v", empty, want)
	}

	// 02_kafka_tables.sql must have received the broker list: a Kafka engine table that
	// cannot reach a broker silently consumes nothing.
	if !brokerRendered {
		t.Errorf("no statement carries the rendered kafka_broker_list")
	}
}

func TestInitScriptsMissingBrokerParamIsAnError(t *testing.T) {
	for _, s := range InitScripts() {
		if !strings.Contains(string(s.Bytes), "KafkaBrokerList") {
			continue
		}
		if _, err := dorisinit.Render(s.Name, s.Bytes, dorisinit.Params{}); err == nil {
			t.Errorf("%s rendered without the KafkaBrokerList parameter; a missing broker list would ingest nothing", s.Name)
		}
	}
}
