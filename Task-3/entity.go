package main

type Job struct {
	ID      string
	JobType string // EX: "notification" or "report_updater" or "index_search_system"
	Data    string // As a JSON string or Text
}
