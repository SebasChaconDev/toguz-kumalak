// Disposable Neon fixture installer. Credentials arrive through inherited env,
// and the operator captures stdout in memory before encrypting it with DPAPI.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

var count int
var records []map[string]any

func fail(label string, err error) {
	code := "connection-or-operator-error"
	if e, ok := err.(*pgconn.PgError); ok {
		code = e.Code
	}
	_ = json.NewEncoder(os.Stdout).Encode(map[string]any{"ok": false, "label": label, "sqlstate": code, "statements": count, "records": records})
	os.Exit(1)
}
func exec(c *pgx.Conn, label, sql string) {
	count++
	_, e := c.Exec(context.Background(), sql)
	if e != nil {
		fail(label, e)
	}
	records = append(records, map[string]any{"ordinal": count, "label": label, "status": "PASS"})
}
func connect(dsn, label string) *pgx.Conn {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c, e := pgx.Connect(ctx, dsn)
	if e != nil {
		fail(label, e)
	}
	return c
}
func main() {
	admin := os.Getenv("SPIKE_ADMIN_DSN")
	u, e := url.Parse(admin)
	if e != nil || !strings.Contains(u.Host, "ep-wandering-snow-b8hkpxdu") {
		fail("owned-project-guard", fmt.Errorf("guard"))
	}
	if os.Getenv("SPIKE_SETUP_APPROVED") != "cold-king-37625726" {
		fail("approval-guard", fmt.Errorf("guard"))
	}
	c := connect(admin, "admin-connect")
	defer c.Close(context.Background())
	for _, name := range []string{"identity", "game"} {
		pw := os.Getenv("SPIKE_" + strings.ToUpper(name) + "_DB_PASSWORD")
		if len(pw) < 40 || strings.ContainsAny(pw, "'\\") {
			fail("password-validation", fmt.Errorf("guard"))
		}
		exec(c, "create-"+name+"-owner", "CREATE ROLE spike_"+name+"_owner NOLOGIN NOSUPERUSER NOCREATEDB NOCREATEROLE")
		exec(c, "create-"+name+"-app", "CREATE ROLE spike_"+name+"_app LOGIN PASSWORD '"+pw+"' NOSUPERUSER NOCREATEDB NOCREATEROLE CONNECTION LIMIT 4")
	}
	exec(c, "admin-owner-membership", "GRANT spike_identity_owner,spike_game_owner TO spike_admin")
	for _, name := range []string{"identity", "game"} {
		exec(c, "create-"+name+"-database", "CREATE DATABASE spike_"+name+" OWNER spike_"+name+"_owner")
		exec(c, "revoke-public-"+name+"-connect", "REVOKE ALL ON DATABASE spike_"+name+" FROM PUBLIC")
		exec(c, "grant-"+name+"-connect", "GRANT CONNECT ON DATABASE spike_"+name+" TO spike_"+name+"_app")
	}
	dsns := map[string]string{}
	for _, name := range []string{"identity", "game"} {
		own := *u
		own.Path = "/spike_" + name
		db := connect(own.String(), "admin-"+name+"-connect")
		exec(db, "revoke-"+name+"-schema-create", "REVOKE CREATE ON SCHEMA public FROM PUBLIC")
		if name == "identity" {
			exec(db, "identity-table", "CREATE TABLE synthetic_identity(id text PRIMARY KEY)")
			exec(db, "identity-synthetic-row", "INSERT INTO synthetic_identity VALUES('alice')")
			exec(db, "identity-minimal-select", "GRANT USAGE ON SCHEMA public TO spike_identity_app")
			exec(db, "identity-minimal-table-select", "GRANT SELECT ON synthetic_identity TO spike_identity_app")
		} else {
			up := []string{"CREATE TABLE synthetic_games(id text PRIMARY KEY,owner text NOT NULL,revision bigint NOT NULL CHECK(revision>=0),effects bigint NOT NULL CHECK(effects>=0))", "CREATE TABLE synthetic_receipts(game_id text REFERENCES synthetic_games(id),actor text NOT NULL,operation_id text NOT NULL,intent_hash text NOT NULL,revision bigint NOT NULL,effects bigint NOT NULL,PRIMARY KEY(game_id,actor,operation_id))", "INSERT INTO synthetic_games VALUES('owned-game','alice',0,0),('foreign-game','bob',0,0)"}
			for i, sql := range up {
				exec(db, fmt.Sprintf("game-up-first-%d", i+1), sql)
			}
			exec(db, "game-down-receipts", "DROP TABLE synthetic_receipts")
			exec(db, "game-down-games", "DROP TABLE synthetic_games")
			for i, sql := range up {
				exec(db, fmt.Sprintf("game-up-reinstall-%d", i+1), sql)
			}
			exec(db, "game-minimal-schema", "GRANT USAGE ON SCHEMA public TO spike_game_app")
			exec(db, "game-minimal-tables", "GRANT SELECT,INSERT,UPDATE ON synthetic_games,synthetic_receipts TO spike_game_app")
		}
		db.Close(context.Background())
		pooled := *u
		pooled.Host = strings.Replace(u.Host, "ep-wandering-snow-b8hkpxdu.", "ep-wandering-snow-b8hkpxdu-pooler.", 1)
		pooled.Path = "/spike_" + name
		pooled.User = url.UserPassword("spike_"+name+"_app", os.Getenv("SPIKE_"+strings.ToUpper(name)+"_DB_PASSWORD"))
		q := pooled.Query()
		q.Set("application_name", "tk-spike01-"+name)
		pooled.RawQuery = q.Encode()
		dsns[name] = pooled.String()
	}
	_ = json.NewEncoder(os.Stdout).Encode(map[string]any{"ok": true, "statements": count, "records": records, "appDSNs": dsns})
}
