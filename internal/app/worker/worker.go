package worker

import (
	"context"
	"sync"

	"github.com/bassga/scraper-bot/internal/domain/downloader"
	"github.com/bassga/scraper-bot/internal/domain/logger"
)

type Worker struct {
	Downloader downloader.Downloader
	Logger logger.Logger
	Ctx context.Context
	Jobs <-chan Job
	Folder string
	MaxRetries int
	RetryStrategy RetryStrategy
	JobHandler JobHandler
}

func (w *Worker) Run(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if w.shouldTerminate() {
			return
		}
		job, ok := w.receiveJob()
		if !ok {
			return
		}
		w.handleJob(job)
	}
}

func (w *Worker) shouldTerminate() bool {
	select {
	case <-w.Ctx.Done():
		// コンテキストキャンセルを検知して終了
		w.Logger.Info("worker received cancellation signal")
		return true
	default: return false
	}
}

func (w *Worker) receiveJob() (Job, bool) {
	job, ok := <-w.Jobs
	return job, ok
}

func (w *Worker) handleJob(job Job) {
	err := w.processJob(job)
	if err != nil {
		w.Logger.Error("giving up downloading image after retries: %v", err)
		return
	}
	w.Logger.Info("successfully downloaded: %s\n", w.Folder)
}

func (w *Worker) processJob(job Job) error {
	return w.RetryStrategy.Do(w.Ctx, job, w.MaxRetries, func(ctx context.Context) error {
		return w.JobHandler.Handle(ctx, job)
	})
}




