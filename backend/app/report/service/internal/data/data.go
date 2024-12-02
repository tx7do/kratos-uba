package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/redis/go-redis/v9"

	"github.com/tx7do/go-utils/entgo"

	conf "github.com/tx7do/kratos-bootstrap/api/gen/go/conf/v1"
	redisClient "github.com/tx7do/kratos-bootstrap/cache/redis"
	bRegistry "github.com/tx7do/kratos-bootstrap/registry"

	"kratos-uba/app/report/service/internal/data/ent"
)

// Data .
type Data struct {
	log *log.Helper

	db  *entgo.EntClient[*ent.Client]
	rdb *redis.Client
}

// NewData .
func NewData(entClient *entgo.EntClient[*ent.Client], redisClient *redis.Client, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data/report-service"))

	d := &Data{
		db:  entClient,
		rdb: redisClient,
		log: l,
	}

	return d, func() {
		l.Info("message", "closing the data resources")
		d.db.Close()
		if err := d.rdb.Close(); err != nil {
			l.Error(err)
		}
	}, nil
}

// NewRedisClient 创建Redis客户端
func NewRedisClient(cfg *conf.Bootstrap, _ log.Logger) *redis.Client {
	//l := log.NewHelper(log.With(logger, "module", "redis/data/admin-service"))
	return redisClient.NewClient(cfg.Data)
}

// NewDiscovery 创建服务发现客户端
func NewDiscovery(cfg *conf.Bootstrap) registry.Discovery {
	return bRegistry.NewDiscovery(cfg.Registry)
}
