package interactor

import (
	"farhadiis/order-process/application/presenter"
	"farhadiis/order-process/application/repository"
	"farhadiis/order-process/domain/model"
	"fmt"
)

type OrderInteractor interface {
	GetAll() ([]*model.Order, error)
	Get(string) (*model.Order, error)
	ListenNewOrder(*model.Order)
	Process(*model.Order) error
	Delete(string) (bool, error)
}

type orderInteractor struct {
	OrderBroker     repository.OrderBroker
	OrderRepository repository.OrderRepository
	OrderPresenter  presenter.OrderPresenter
}

func (t *orderInteractor) GetAll() ([]*model.Order, error) {
	panic("Unimplemented...")
}

func (t *orderInteractor) Get(id string) (*model.Order, error) {
	panic("Unimplemented...")
}

func (t *orderInteractor) ListenNewOrder(order *model.Order) {
	fmt.Println("new order received", order)
	err := t.Process(order)
	if err != nil {
		fmt.Println("failed order process", order)
		return
	}
}

func (t *orderInteractor) Process(order *model.Order) error {
	err := t.OrderRepository.InsertOne(order)
	if err != nil {
		return err
	}
	return nil
}

func (t *orderInteractor) Delete(id string) (bool, error) {
	panic("Unimplemented...")
}

func NewOrderInteractor(r repository.OrderBroker, t repository.OrderRepository, p presenter.OrderPresenter) OrderInteractor {
	l := orderInteractor{r, t, p}
	r.Subscribe(l.ListenNewOrder)
	return &l
}
