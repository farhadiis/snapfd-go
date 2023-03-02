package interactor

import (
	"farhadiis/order/application/presenter"
	"farhadiis/order/application/repository"
	"farhadiis/order/domain/model"
	"time"
)

type OrderInteractor interface {
	GetAll() ([]*model.Order, error)
	Get(string) (*model.Order, error)
	Create(*model.Order) error
	Delete(string) (bool, error)
}

type orderInteractor struct {
	OrderBroker    repository.OrderBroker
	OrderPresenter presenter.OrderPresenter
}

func (t *orderInteractor) GetAll() ([]*model.Order, error) {
	panic("Unimplemented...")
}

func (t *orderInteractor) Get(id string) (*model.Order, error) {
	panic("Unimplemented...")
}

func (t *orderInteractor) Create(order *model.Order) error {
	order.CreatedAt = time.Now()
	err := t.OrderBroker.Publish(order)
	if err != nil {
		return err
	}
	return nil
}

func (t *orderInteractor) Delete(id string) (bool, error) {
	panic("Unimplemented...")
}

func NewOrderInteractor(r repository.OrderBroker, p presenter.OrderPresenter) OrderInteractor {
	return &orderInteractor{r, p}
}
