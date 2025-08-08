package main

import (
	"math/rand/v2"
	"time"

	"github.com/avast/retry-go"
)

const percentOfFields = 80.0
const maxRetries = 3
const minDelay = 100 * time.Millisecond
const maxDelay = 1 * time.Second

func randomException() bool {
	return rand.Float64() >= percentOfFields/100
}

func retryProgress(fn func() error) error {
	return retry.Do(
		fn,
		retry.Attempts(maxRetries),
		retry.Delay(minDelay),
		retry.MaxDelay(maxDelay),
	)
}
