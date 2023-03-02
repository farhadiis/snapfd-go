package registry

import (
	"database/sql"
	"farhadiis/order-process/interface/controller"
	"github.com/redis/go-redis/v9"
)

type Registry interface {
	NewAppController() controller.AppController
}

type registry struct {
	RedisClient *redis.Client
	MySQLClient *sql.DB
}

func NewRegistry(redisClient *redis.Client, mySqlClient *sql.DB) Registry {
	return &registry{redisClient, mySqlClient}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewOrderController()
}
