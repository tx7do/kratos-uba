// Package clickhouseinit turns the ClickHouse provisioning scripts in backend/sql/clickhouse
// into an idempotent, one-shot apply: ensure the database, cut each script into statements,
// render its parameters, and execute them in order.
//
// Statement lexing and template rendering are shared with dorisinit — quoting, comments and
// ';' splitting are SQL-generic. This package only adds what ClickHouse needs: a native
// connection surface and database creation. There are no jobs to supervise, unlike Doris's
// Routine Load: the Kafka engine tables themselves are the consumers, so "apply the schema"
// is the whole provisioning story.
package clickhouseinit

import (
	"context"
	"fmt"
	"strings"

	"go-wind-uba/pkg/dorisinit"
)

// Script is one provisioning script; Params are its template parameters. Both are the
// dorisinit shapes, so the ingest CLI can pass one script list type around.
type Script = dorisinit.Script
type Params = dorisinit.Params

// Statement is one top-level SQL statement, as cut by dorisinit.SplitStatements.
type Statement = dorisinit.Statement

// Conn is the execution surface. *clickhouse.Conn (clickhouse-go v2) satisfies it.
type Conn interface {
	Exec(ctx context.Context, query string, args ...any) error
	Close() error
}

// ApplyOptions controls one apply run.
type ApplyOptions struct {
	// Database is ensured to exist before any script runs; empty skips it. The scripts
	// themselves address gw_uba explicitly, so this should stay gw_uba unless they are
	// edited to match.
	Database string
	// DryRun prints every statement without sending anything (the database is not created).
	DryRun bool
}

// Result records one statement a run visited.
type Result struct {
	Script string
	Line   int
	Title  string
	Err    error
}

// Apply executes the scripts in order on one connection. Every statement is idempotent
// (CREATE ... IF NOT EXISTS), so a re-run only fills what a previous run stopped short of.
// It stops at the first failing statement and reports what already landed.
func Apply(ctx context.Context, conn Conn, scripts []Script, params Params, opts ApplyOptions) ([]Result, error) {
	var results []Result

	if opts.Database != "" && !opts.DryRun {
		ddl := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", opts.Database)
		err := conn.Exec(ctx, ddl)
		results = append(results, Result{Script: "<database>", Title: ddl, Err: err})
		if err != nil {
			return results, fmt.Errorf("create database %s: %w", opts.Database, err)
		}
	}

	for _, script := range scripts {
		src, err := dorisinit.Render(script.Name, script.Bytes, params)
		if err != nil {
			return results, fmt.Errorf("%s: %w", script.Name, err)
		}
		stmts, err := dorisinit.SplitStatements(dorisinit.StripBOM(src))
		if err != nil {
			return results, fmt.Errorf("%s: %w", script.Name, err)
		}

		for _, stmt := range stmts {
			result := Result{Script: script.Name, Line: stmt.Line, Title: Abbreviate(stmt.Code)}
			if !opts.DryRun {
				if err := conn.Exec(ctx, stmt.Text); err != nil {
					result.Err = err
					results = append(results, result)
					return results, fmt.Errorf("%s:%d: %w", script.Name, stmt.Line, err)
				}
			}
			results = append(results, result)
		}
	}
	return results, nil
}

// Abbreviate collapses a statement to one short line for the run report.
func Abbreviate(s string) string {
	s = strings.Join(strings.Fields(s), " ")
	runes := []rune(s)
	if len(runes) > 72 {
		return string(runes[:72]) + "…"
	}
	return s
}
