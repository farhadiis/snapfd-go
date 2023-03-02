package repository

import (
	"database/sql"
	"farhadiis/order-process/application/repository"
	"farhadiis/order-process/domain/model"
	"fmt"
)

type orderRepository struct {
	db *sql.DB
}

func (t *orderRepository) FindAll() ([]*model.Order, error) {
	panic("Unimplemented...")
}

func (t *orderRepository) FindOne(id string) (*model.Order, error) {
	panic("Unimplemented...")
}

func (t *orderRepository) InsertOne(order *model.Order) error {
	_, err := t.db.Exec("INSERT INTO orders (order_id, title, price) VALUES (?, ?, ?)", order.OrderId, order.Title, order.Price)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (t *orderRepository) DeleteOne(id string) (int64, error) {
	panic("Unimplemented...")
}

func (t *orderRepository) initialize() {
	_, err := t.db.Exec(`
		CREATE TABLE IF NOT EXISTS orders (
		  order_id   INT NOT NULL PRIMARY KEY,
		  title      VARCHAR(128) NOT NULL,
		  price      FLOAT NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
}

func NewOrderRepository(client *sql.DB) repository.OrderRepository {
	r := orderRepository{client}
	r.initialize()
	return &r
}
