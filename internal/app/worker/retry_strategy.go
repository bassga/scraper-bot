package worker

import (
	"context"
	"time"
)

type RetryStrategy interface {
	Do(ctx context.Context, job Job, maxRetries int, action func(context.Context) error) error
}

// 特定の時間sleepしてretry
type FixedRetryStrategy struct{}
func (f *FixedRetryStrategy) Do(ctx context.Context, job Job, maxRetries int, action func(context.Context) error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		attemptCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		err = action(attemptCtx)
		if err == nil {
			return nil // 成功したらすぐreturn
		}
		time.Sleep(1 * time.Second)
	}
	return err
}

// retryしない
type NoRetryStrategy struct{}
func (n *NoRetryStrategy) Do(ctx context.Context, job Job, maxRetries int, action func(context.Context) error) error {
	ctxWithTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return action(ctxWithTimeout)
}

// 指数関数的にsleep時間を増やす
type ExponentialBackoffStrategy struct{}
func (e *ExponentialBackoffStrategy) Do(ctx context.Context, job Job, maxRetries int, action func(context.Context) error) error {
	var err error
	for i := 0; i < maxRetries; i++ {
		attemptCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		err = action(attemptCtx)
		if err == nil {
			return nil // 成功したらすぐreturn
		}
		sleep := time.Duration(1<<i) * time.Second // 1s, 2s, 4s, 8s...
		time.Sleep(sleep)
	}
	return err
}