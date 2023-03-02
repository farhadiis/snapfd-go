package interactor

import (
	"farhadiis/order-process/application/interactor"
	"farhadiis/order-process/application/repository"
	"farhadiis/order-process/domain/model"
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
	MockPresenter  struct{}
	MockRepository struct {
		Orders []*model.Order
	}
	MockBroker struct {
		Order *model.Order
	}
)

func (t *MockRepository) FindAll() ([]*model.Order, error) {
	panic("Unimplemented...")
}

func (t *MockRepository) FindOne(id string) (*model.Order, error) {
	panic("Unimplemented...")
}

func (t *MockRepository) InsertOne(order *model.Order) error {
	if order == nil {
		return nil
	}
	t.Orders = append(t.Orders, order)
	return nil
}

func (t *MockRepository) DeleteOne(id string) (int64, error) {
	panic("Unimplemented...")
}

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

// TestListenNewOrder
func TestListenNewOrder(t *testing.T) {
	mockBroker := &MockBroker{}
	mockRepository := &MockRepository{}
	mockPresenter := &MockPresenter{}
	in := interactor.NewOrderInteractor(mockBroker, mockRepository, mockPresenter)
	in.ListenNewOrder(mockOrder)
	if mockRepository.Orders[0] != mockOrder {
		t.Fatalf(`Get() must return %q, nill`, mockOrder)
	}
}
