package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	brokerService := NewBrokerService()
	brokerService.CreateQueue("update_profile")

	service := NewServiceUpdateProfile(brokerService)

	for i := 0; i < 100_000; i++ {
		job := Job{
			ID:         fmt.Sprintf("job-%d", i),
			Data:       fmt.Sprintf("data-%d", i),
			retryCount: 0,
			maxRetries: rand.Intn(4),
		}
		brokerService.AppendToQueue("update_profile", job)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 40*time.Second)
	defer cancel()

	if err := service.Run(ctx); err != nil {
		log.Fatalf("Service failed: %v", err)
	}
}
