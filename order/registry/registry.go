package registry

import (
	"farhadiis/order/interface/controller"
	"github.com/redis/go-redis/v9"
)

type Registry interface {
	NewAppController() controller.AppController
}

type registry struct {
	RedisClient *redis.Client
}

func NewRegistry(redisClient *redis.Client) Registry {
	return &registry{redisClient}
}

func (r *registry) NewAppController() controller.AppController {
	return r.NewOrderController()
}
