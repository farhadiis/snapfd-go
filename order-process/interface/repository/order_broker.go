package repository

import (
	"context"
	"encoding/json"
	"farhadiis/order-process/application/repository"
	"farhadiis/order-process/domain/model"
	"github.com/redis/go-redis/v9"
	"log"
)

type orderBroker struct {
	client *redis.Client
}

var channel = "orders"

func (t *orderBroker) Publish(order *model.Order) error {
	err := t.client.Publish(context.Background(), channel, order).Err()
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (t *orderBroker) Subscribe(handler repository.SubscribeHandler) {
	subscriber := t.client.Subscribe(context.Background(), channel)
	ch := subscriber.Channel()
	go func() {
		for msg := range ch {
			if handler != nil {
				bytes := []byte(msg.Payload)
				var order *model.Order
				err := json.Unmarshal(bytes, &order)
				if err != nil {
					return
				}
				handler(order)
			}
		}
	}()
}

func NewOrderBroker(client *redis.Client) repository.OrderBroker {
	return &orderBroker{client}
}
