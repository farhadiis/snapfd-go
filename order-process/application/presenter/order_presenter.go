package presenter

import "farhadiis/order-process/domain/model"

type OrderPresenter interface {
	ResponseOrders([]*model.Order) []*model.Order
}
