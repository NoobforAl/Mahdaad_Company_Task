package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

type ServiceUpdateProfile struct {
	BrokerService *BrokerService

	isUp bool

	lastCheckedAt time.Time
}

func NewServiceUpdateProfile(brokerService *BrokerService) *ServiceUpdateProfile {
	return &ServiceUpdateProfile{
		BrokerService: brokerService,
		isUp:          true,
		lastCheckedAt: time.Now(),
	}
}

func (s *ServiceUpdateProfile) CheckStatus() bool {
	return s.lastCheckedAt.Add(time.Second).After(time.Now())
}

func (s *ServiceUpdateProfile) processJob(job Job) error {
	do := func() error {
		if ok := randomException(); ok {
			return fmt.Errorf("failed to process job %v", job)
		}
		return nil
	}

	return retryProgress(do)
}

func (s *ServiceUpdateProfile) Run(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()

		default:
			if !s.isUp {
				if ok := s.CheckStatus(); !ok {
					s.isUp = true
				}
			}

			job, ok := s.BrokerService.PopFromQueue("update_profile")
			if !ok {
				log.Println("No jobs in the queue, waiting...")
				time.Sleep(1 * time.Second)
				continue
			}

			log.Println("Processing job:", job)
			if err := s.processJob(job); err != nil {
				s.isUp = false
				log.Printf("Error processing job %v: %v", job, err)

				job.retryCount++
				if job.retryCount > job.maxRetries {
					log.Printf("Max retries reached for job %v", job)
					continue
				}

				s.BrokerService.AppendToQueue("update_profile", job)
				continue
			}

			log.Println("Job processed successfully:", job)
		}
	}
}
