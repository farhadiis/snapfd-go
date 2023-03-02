package repository

import "farhadiis/order/domain/model"

type SubscribeHandler func(order *model.Order)

type OrderBroker interface {
	Publish(*model.Order) error
	Subscribe(handler SubscribeHandler)
}
