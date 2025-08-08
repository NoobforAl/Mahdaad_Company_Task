package main

type Item struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

type Order struct {
	ID    string
	Items []*Item
}

type Payment struct {
	ID      string
	OrderID string
	Amount  float64
	Status  string
}

type Job struct {
	ID      string
	Order   *Order
	Payment *Payment
}
