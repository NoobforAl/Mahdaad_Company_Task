package main

import (
	"container/list"
)

type BrokerService struct {
	queues map[string]*list.List
}

func NewBrokerService() *BrokerService {
	return &BrokerService{
		queues: make(map[string]*list.List),
	}
}

func (bs *BrokerService) CreateQueue(queueName string) {
	bs.queues[queueName] = list.New()
}

func (bs *BrokerService) AppendToQueue(queueName string, job Job) {
	if queue, exists := bs.queues[queueName]; exists {
		queue.PushBack(job)
	}
}

func (bs *BrokerService) PopFromQueue(queueName string) (Job, bool) {
	if queue, exists := bs.queues[queueName]; exists {
		if front := queue.Front(); front != nil {
			queue.Remove(front)
			return front.Value.(Job), true
		}
	}
	return Job{}, false
}
