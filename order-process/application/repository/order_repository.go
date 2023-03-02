package repository

import "farhadiis/order-process/domain/model"

type OrderRepository interface {
	FindAll() ([]*model.Order, error)
	FindOne(string) (*model.Order, error)
	InsertOne(*model.Order) error
	DeleteOne(string) (int64, error)
}
