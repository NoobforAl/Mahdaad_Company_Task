package main

import "fmt"

type PaymentService struct {
	store *StoreService

	payments map[string]*Payment
}

func NewPaymentService(store *StoreService) *PaymentService {
	return &PaymentService{
		store:    store,
		payments: make(map[string]*Payment),
	}
}

func (ps *PaymentService) DoPayment(payment *Payment) bool {
	return !randomException()
}

func (ps *PaymentService) NewTask(job Job) (Task, error) {
	amount := 0.0
	for _, item := range job.Order.Items {
		amount += item.Price * float64(item.Quantity)
	}

	payment := &Payment{
		ID:      "payment",
		OrderID: job.Order.ID,
		Amount:  amount,
		Status:  "pending",
	}

	job.Payment = payment

	return &PaymentTask{
		Job:   job,
		store: ps.store,
	}, nil
}

type PaymentTask struct {
	Job            Job
	store          *StoreService
	paymentService *PaymentService
}

func (pt *PaymentTask) Execute() error {
	if !pt.paymentService.DoPayment(pt.Job.Payment) {
		return fmt.Errorf("payment failed")
	}
	return nil
}

func (pt *PaymentTask) Rollback() error {
	return nil
}
