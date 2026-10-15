package clickhouseinit

import (
	"context"
	"strings"
	"testing"

	"go-wind-uba/pkg/dorisinit"
)

// fakeConn records every statement instead of sending it, so tests can assert on the plan.
type fakeConn struct {
	sent []string
	err  map[string]error // statement prefix → failure
}

func (c *fakeConn) Exec(_ context.Context, query string, _ ...any) error {
	for prefix, err := range c.err {
		if strings.HasPrefix(query, prefix) {
			return err
		}
	}
	c.sent = append(c.sent, query)
	return nil
}

func (c *fakeConn) Close() error { return nil }

var baseScript = dorisinit.Script{
	Name: "1_base.sql",
	Bytes: []byte(`-- comment; with a semicolon
CREATE TABLE IF NOT EXISTS gw_uba.events_fact (id UInt32);
CREATE TABLE IF NOT EXISTS gw_uba.sessions_fact (id UInt32);
`),
}

var kafkaScript = dorisinit.Script{
	Name: "02_kafka.sql",
	Bytes: []byte("CREATE TABLE IF NOT EXISTS gw_uba.kafka_raw (raw String)\n" +
		"ENGINE = Kafka SETTINGS kafka_broker_list = '{{.KafkaBrokerList}}';\n"),
}

func params() Params { return Params{"KafkaBrokerList": "kafka:9092"} }

func TestApplyCreatesDatabaseThenScripts(t *testing.T) {
	conn := &fakeConn{}
	results, err := Apply(context.Background(), conn, []Script{baseScript, kafkaScript}, params(),
		ApplyOptions{Database: "gw_uba"})
	if err != nil {
		t.Fatalf("apply: %v", err)
	}
	if len(conn.sent) == 0 || conn.sent[0] != "CREATE DATABASE IF NOT EXISTS gw_uba" {
		t.Fatalf("first statement should create the database, got %q", first(conn.sent))
	}
	if got, want := len(conn.sent), 4; got != want {
		t.Fatalf("sent %d statements, want %d: %q", got, want, conn.sent)
	}
	if !strings.Contains(conn.sent[len(conn.sent)-1], "kafka:9092") {
		t.Fatalf("broker list was not rendered: %q", conn.sent[len(conn.sent)-1])
	}
	if len(results) != len(conn.sent) {
		t.Fatalf("report has %d rows for %d statements", len(results), len(conn.sent))
	}
}

func TestApplyDryRunSendsNothing(t *testing.T) {
	conn := &fakeConn{}
	results, err := Apply(context.Background(), conn, []Script{baseScript, kafkaScript}, params(),
		ApplyOptions{Database: "gw_uba", DryRun: true})
	if err != nil {
		t.Fatalf("apply: %v", err)
	}
	if len(conn.sent) != 0 {
		t.Fatalf("dry run sent %d statements", len(conn.sent))
	}
	if len(results) != 3 {
		t.Fatalf("plan should still list every statement, got %d", len(results))
	}
}

func TestApplySkipsDatabaseWhenUnset(t *testing.T) {
	conn := &fakeConn{}
	if _, err := Apply(context.Background(), conn, []Script{baseScript}, nil, ApplyOptions{}); err != nil {
		t.Fatalf("apply: %v", err)
	}
	if strings.HasPrefix(conn.sent[0], "CREATE DATABASE") {
		t.Fatalf("no database configured, but one was created: %q", conn.sent[0])
	}
}

func TestApplyStopsAtFirstFailureAndReports(t *testing.T) {
	conn := &fakeConn{err: map[string]error{
		"CREATE TABLE IF NOT EXISTS gw_uba.sessions_fact": context.DeadlineExceeded,
	}}
	results, err := Apply(context.Background(), conn, []Script{baseScript}, nil, ApplyOptions{Database: "gw_uba"})
	if err == nil {
		t.Fatal("apply should fail when a statement fails")
	}
	if !strings.Contains(err.Error(), "1_base.sql:3") {
		t.Fatalf("error should name script and line, got: %v", err)
	}
	last := results[len(results)-1]
	if last.Err == nil || last.Line != 3 {
		t.Fatalf("last report row should carry the failure, got %+v", last)
	}
	// The database and the first table landed before the failure; a re-run repairs it.
	if len(conn.sent) != 2 {
		t.Fatalf("statements before the failure should have landed, got %q", conn.sent)
	}
}

func TestApplyMissingParameterIsAnError(t *testing.T) {
	conn := &fakeConn{}
	if _, err := Apply(context.Background(), conn, []Script{kafkaScript}, Params{}, ApplyOptions{}); err == nil {
		t.Fatal("rendering without the broker list parameter should fail")
	}
	if len(conn.sent) != 0 {
		t.Fatalf("nothing should be sent when rendering fails, got %q", conn.sent)
	}
}

func TestApplySemicolonInsideCommentDoesNotSplit(t *testing.T) {
	conn := &fakeConn{}
	if _, err := Apply(context.Background(), conn, []Script{baseScript}, nil, ApplyOptions{}); err != nil {
		t.Fatalf("apply: %v", err)
	}
	for _, s := range conn.sent {
		if strings.Contains(s, "comment") {
			t.Fatalf("comment text leaked into a statement: %q", s)
		}
	}
}

func first(ss []string) string {
	if len(ss) == 0 {
		return ""
	}
	return ss[0]
}
