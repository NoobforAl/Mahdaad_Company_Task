package main

type Job struct {
	ID      string
	JobType string
	Data    string

	retryCount int // Number of times the job has been retried
	maxRetries int // Maximum number of retries allowed for the job
}
