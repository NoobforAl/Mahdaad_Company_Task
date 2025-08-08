package main

import (
	"fmt"
	"sync"
)

type OrderService struct {
	Orders map[string]Order

	store *StoreService

	mu sync.Mutex
}

func NewOrderService(store *StoreService) *OrderService {
	return &OrderService{
		Orders: make(map[string]Order),
		store:  store,
	}
}

func (s *OrderService) NewTask(job Job) (Task, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &OrderTask{
		Job:   job,
		store: s.store,
	}, nil
}

type OrderTask struct {
	Job

	store *StoreService
}

func (o *OrderTask) Execute() error {
	for _, item := range o.Order.Items {
		_, ok := o.store.GetItem(item.ID)
		if !ok {
			return fmt.Errorf("item not found in store: %s", item.ID)
		}
	}

	if ok := randomException(); ok {
		return fmt.Errorf("order failed validation")
	}

	return nil
}

func (o *OrderTask) Rollback() error {
	return nil
}
