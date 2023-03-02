package registry

import (
	"farhadiis/order-process/application/interactor"
	p "farhadiis/order-process/application/presenter"
	r "farhadiis/order-process/application/repository"
	"farhadiis/order-process/interface/controller"
	"farhadiis/order-process/interface/presenter"
	"farhadiis/order-process/interface/repository"
)

func (r *registry) NewOrderController() controller.OrderController {
	return controller.NewOrderController(r.NewOrderInteractor())
}

func (r *registry) NewOrderInteractor() interactor.OrderInteractor {
	return interactor.NewOrderInteractor(r.NewOrderBroker(), r.NewOrderRepository(), r.NewOrderPresenter())
}

func (r *registry) NewOrderBroker() r.OrderBroker {
	return repository.NewOrderBroker(r.RedisClient)
}

func (r *registry) NewOrderRepository() r.OrderRepository {
	return repository.NewOrderRepository(r.MySQLClient)
}

func (r *registry) NewOrderPresenter() p.OrderPresenter {
	return presenter.NewOrderPresenter()
}
