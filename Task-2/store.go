package main

import (
	"fmt"
	"sync"
)

type StoreService struct {
	Inventory map[string]Item

	mu sync.RWMutex
}

func NewStoreService() *StoreService {
	store := &StoreService{
		Inventory: make(map[string]Item),
	}

	store.Inventory["item1"] = Item{
		ID:       "item1",
		Name:     "Item 1",
		Price:    10.0,
		Quantity: 10,
	}
	store.Inventory["item2"] = Item{
		ID:       "item2",
		Name:     "Item 2",
		Price:    5.0,
		Quantity: 5,
	}
	store.Inventory["item3"] = Item{
		ID:       "item3",
		Name:     "Item 3",
		Price:    3.0,
		Quantity: 3,
	}

	return store
}

func (s *StoreService) GetItem(id string) (*Item, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	item, exists := s.Inventory[id]
	return &item, exists
}

func (s *StoreService) increaseItemQuantity(id string, amount int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.Inventory[id]
	if !exists {
		return false
	}

	item.Quantity += amount
	s.Inventory[id] = item
	return true
}

func (s *StoreService) decreaseItemQuantity(id string, amount int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	item, exists := s.Inventory[id]
	if !exists {
		return false
	}

	if item.Quantity < amount {
		return false
	}

	item.Quantity -= amount
	s.Inventory[id] = item
	return true
}

func (s *StoreService) NewTask(job Job) (Task, error) {
	task := &StoreTask{
		store: s,
		Job:   job,
	}
	return task, nil
}

type StoreTask struct {
	Job

	store *StoreService

	decreasedItems []*Item
}

func (s *StoreTask) Execute() error {
	// Sample TRX ( better to SQL like TRX)
	for _, item := range s.Order.Items {
		if !s.store.decreaseItemQuantity(item.ID, item.Quantity) {
			return fmt.Errorf("failed to process item: %s", item.ID)
		}
		s.decreasedItems = append(s.decreasedItems, item)
	}

	if ok := randomException(); ok {
		return fmt.Errorf("store failed validation")
	}

	return nil
}

func (s *StoreTask) Rollback() error {
	for _, item := range s.decreasedItems {
		if !s.store.increaseItemQuantity(item.ID, item.Quantity) {
			return fmt.Errorf("failed to rollback item: %s", item.ID)
		}
	}
	return nil
}
