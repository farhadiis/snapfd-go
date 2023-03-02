package interactor

import (
	"farhadiis/order/application/interactor"
	"farhadiis/order/application/repository"
	"farhadiis/order/domain/model"
	"testing"
	"time"
)

var (
	mockOrder = &model.Order{
		OrderId:   10,
		Title:     "burger",
		Price:     1000,
		CreatedAt: time.Time{},
	}
)

type (
	MockPresenter struct{}
	MockBroker    struct {
		Order *model.Order
	}
)

func (m *MockBroker) Publish(order *model.Order) error {
	m.Order = order
	return nil
}

func (m *MockBroker) Subscribe(handler repository.SubscribeHandler) {
	handler(m.Order)
}

func (m *MockPresenter) ResponseOrders(orders []*model.Order) []*model.Order {
	return orders
}

// TestPublish calls interactor.Create, then checking broker to exist in queue.
func TestPublish(t *testing.T) {
	mockBroker := &MockBroker{}
	mockPresenter := &MockPresenter{}
	in := interactor.NewOrderInteractor(mockBroker, mockPresenter)
	err := in.Create(mockOrder)
	if err != nil {
		t.Fatalf("Can't create order")
	}
	mockBroker.Subscribe(func(order *model.Order) {
		if order != mockOrder {
			t.Fatalf(`Get() must return %q, nill`, mockOrder)
		}
	})
}
