package data

import (
	"kratos-bi/pkg/bootstrap"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v4/stdlib"
	_ "github.com/lib/pq"

	"kratos-bi/app/logger/service/internal/data/ent"
	"kratos-bi/gen/api/go/common/conf"
)

// Data .
type Data struct {
	log *log.Helper

	db  *ent.Client
	rdb *redis.Client
}

// NewData .
func NewData(entClient *ent.Client, redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/logger-service"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		if err := d.db.Close(); err != nil {
			l.Error(err)
		}
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, logger log.Logger) *redis.Client {
	l := log.NewHelper(log.With(logger, "module", "redis/data/logger-service"))
	return bootstrap.NewRedisClient(cfg, l)
}

// NewEntClient 创建数据库客户端
func NewEntClient(cfg *conf.Bootstrap, logger log.Logger) *ent.Client {
	l := log.NewHelper(log.With(logger, "module", "ent/data/logger-service"))

	client, err := ent.Open(
		cfg.Data.Database.Driver,
		cfg.Data.Database.Source,
	)
	if err != nil {
		l.Fatalf("failed opening connection to db: %v", err)
	}

	return client
}
