package main

import "log"

func rollback(tasks []Task) {
	for i := len(tasks) - 1; i >= 0; i-- {
		task := tasks[i]

		if err := retryProgress(task.Rollback); err != nil {
			log.Println("This never happened!")
		}
	}
}

func DoSomeJob() error {
	storeService := NewStoreService()
	orderService := NewOrderService(storeService)
	paymentService := NewPaymentService(storeService)

	myOrder := Order{
		ID: "order1",
		Items: []*Item{
			{ID: "item1", Name: "Item 1", Price: 10.0, Quantity: 1},
			{ID: "item2", Name: "Item 2", Price: 5.0, Quantity: 2},
			{ID: "item3", Name: "Item 3", Price: 3.0, Quantity: 3},
		},
	}

	job := Job{
		ID:      "job1",
		Order:   &myOrder,
		Payment: nil,
	} // set a method to set some values

	var tasks = make([]Task, 0, 3)
	orderTask, err := orderService.NewTask(job)
	if err != nil {
		return err
	}

	tasks = append(tasks, orderTask)
	if err := orderTask.Execute(); err != nil {
		defer rollback(tasks)
		return err
	}

	storeTask, err := storeService.NewTask(job)
	if err != nil {
		log.Fatalf("failed to create store task: %v", err)
		defer rollback(tasks)
		return err
	}

	tasks = append(tasks, storeTask)
	if err := storeTask.Execute(); err != nil {
		defer rollback(tasks)
		return err
	}

	paymentTask, err := paymentService.NewTask(job)
	if err != nil {
		log.Fatalf("failed to create payment task: %v", err)
		defer rollback(tasks)
		return err
	}

	tasks = append(tasks, paymentTask)
	if err := paymentTask.Execute(); err != nil {
		defer rollback(tasks)
		return err
	}

	log.Println("All tasks executed successfully")
	return nil
}

func main() {
	// try 10 times
	for range 10 {
		if err := DoSomeJob(); err != nil {
			log.Printf("Job failed: %v", err)
		}
	}
}
