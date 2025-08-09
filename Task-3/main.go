package main

import (
	"log"
	"sync"
)

func main() {
	broker := NewBrokerService()
	broker.RegisterQueue("notification", 10)
	broker.RegisterQueue("report_updater", 10)
	broker.RegisterQueue("index_search_system", 10)

	jobs := []Job{
		{
			ID:      "1",
			JobType: "notification",
			Data:    `{"message": "Hello, World!"}`,
		},
		{
			ID:      "2",
			JobType: "report_updater",
			Data:    `{"report_id": "12345"}`,
		},
		{
			ID:      "3",
			JobType: "index_search_system",
			Data:    `{"index": "user_profiles"}`,
		},
	}

	for _, j := range jobs {
		err := broker.Publish(j)
		if err != nil {
			log.Fatal(err)
		}
	}

	var wg sync.WaitGroup
	wg.Add(3)

	notificationSub := broker.Subscribe("notification")
	go func() {
		for job := range notificationSub {
			log.Printf("Notification Received job: %+v\n", job)
			wg.Done()
		}
	}()

	reportSub := broker.Subscribe("report_updater")
	go func() {
		for job := range reportSub {
			log.Printf("Report Updater Received job: %+v\n", job)
			wg.Done()
		}
	}()

	indexSub := broker.Subscribe("index_search_system")
	go func() {
		for job := range indexSub {
			log.Printf("Index Search System Received job: %+v\n", job)
			wg.Done()
		}
	}()

	wg.Wait()
}
