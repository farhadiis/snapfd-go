package registry

import (
	"farhadiis/order/application/interactor"
	p "farhadiis/order/application/presenter"
	r "farhadiis/order/application/repository"
	"farhadiis/order/interface/controller"
	"farhadiis/order/interface/presenter"
	"farhadiis/order/interface/repository"
)

func (r *registry) NewOrderController() controller.OrderController {
	return controller.NewOrderController(r.NewOrderInteractor())
}

func (r *registry) NewOrderInteractor() interactor.OrderInteractor {
	return interactor.NewOrderInteractor(r.NewOrderBroker(), r.NewOrderPresenter())
}

func (r *registry) NewOrderBroker() r.OrderBroker {
	return repository.NewOrderBroker(r.RedisClient)
}

func (r *registry) NewOrderPresenter() p.OrderPresenter {
	return presenter.NewOrderPresenter()
}
