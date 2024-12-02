package data

import (
	"database/sql"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"

	_ "github.com/ClickHouse/clickhouse-go/v2"

	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
)

// Data .
type Data struct {
	log *log.Helper

	rdb *redis.Client
	db  *sql.DB
}

// NewData .
func NewData(db *sql.DB, rdb *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/logger-service"))

	d := &Data{
		rdb: rdb,
		db:  db,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")

		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}

		if err := d.db.Close(); err != nil {
			l.Error(err)
		}

	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	//l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))
	return redisClient.NewClient(cfg.Data)
}

// NewClickHouseClient 创建数据库客户端
func NewClickHouseClient(cfg *conf.Bootstrap, logger log.Logger) *sql.DB {
	l := log.NewHelper(log.With(logger, "module", "ent/data/logger-service"))

	conn, err := sql.Open(cfg.Data.Database.Driver, cfg.Data.Database.Source)
	if err != nil {
		l.Errorf("create clickhouse connection failed: %s", err.Error())
		return nil
	}

	if err := conn.Ping(); err != nil {
		l.Errorf("ping clickhouse failed: %s", err.Error())
		return nil
	}

	return conn
}
