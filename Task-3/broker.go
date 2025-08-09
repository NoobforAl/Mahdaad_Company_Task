package main

import (
	"fmt"
	"log"
)

type BrokerService struct {
	queues map[string]chan Job
}

func NewBrokerService() *BrokerService {
	return &BrokerService{
		queues: make(map[string]chan Job),
	}
}

func (b *BrokerService) RegisterQueue(name string, bufferSize int) {
	b.queues[name] = make(chan Job, bufferSize)
}

func (b *BrokerService) Publish(job Job) error {
	ch, ok := b.queues[string(job.JobType)]
	if !ok {
		return fmt.Errorf("queue not found")
	}
	select {
	case ch <- job:
		log.Println("Job published successfully")
	default:
		return fmt.Errorf("queue is full")
	}

	return nil
}

func (b *BrokerService) Subscribe(queueName string) <-chan Job {
	if ch, ok := b.queues[queueName]; ok {
		return ch
	}
	return nil
}
