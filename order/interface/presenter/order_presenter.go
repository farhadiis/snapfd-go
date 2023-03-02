package presenter

import (
	"farhadiis/order/application/presenter"
	"farhadiis/order/domain/model"
)

type orderPresenter struct{}

func (tp *orderPresenter) ResponseOrders(orders []*model.Order) []*model.Order {
	for _, order := range orders {
		order.Title = "# " + order.Title
	}
	return orders
}

func NewOrderPresenter() presenter.OrderPresenter {
	return &orderPresenter{}
}
