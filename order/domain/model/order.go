package model

import (
	"encoding/json"
	"fmt"
	"time"
)

type Order struct {
	OrderId   int64     `json:"order_id"`
	Price     float64   `json:"price" binding:"required"`
	Title     string    `json:"title" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
}

func (t Order) String() string {
	return fmt.Sprintf("#%v (%v, price=%v)", t.OrderId, t.Title, t.Price)
}

func (t Order) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}
