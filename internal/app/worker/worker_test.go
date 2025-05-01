package worker

import (
	"context"
	"fmt"
	"sync"
	"testing"
)


type mockDownloader struct{}

func (m *mockDownloader) DownloadImage(ctx context.Context, url, destFolder, saveAsName string) (string, error) {
	return "dummy/path", nil
}

type mockLogger struct{}

func (m *mockLogger) Info(format string, args ...interface{}) {}
func (m *mockLogger) Error(format string, args ...interface{}) {}

func TestWorker_Run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobChan := make(chan Job, 1)
	jobChan <- Job{URL: "https://example.com/image.png", SaveAsName: "image.png"}

	close(jobChan)
	worker := Worker{
		Downloader: &mockDownloader{},
		Logger: &mockLogger{},
		Ctx: ctx,
		Jobs: jobChan,
		Folder: "dummy_folder",
		MaxRetries: 3,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go worker.Run(&wg)
	wg.Wait()
}


type failingDownloader struct{
	failCount int
}

func (m *failingDownloader) DownloadImage(ctx context.Context, url, destFolder, saveAsName string) (string, error) {
	m.failCount++
	return "", fmt.Errorf("forced error")
}


func TestWorker_Retry(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	jobChan := make(chan Job, 1)
	jobChan <- Job{URL: "https://example.com/image.png", SaveAsName: "image.png"}

	close(jobChan)

	worker := Worker{
		Downloader: &failingDownloader{},
		Logger: &mockLogger{},
		Ctx: ctx,
		Jobs: jobChan,
		Folder: "dummy_folder",
		MaxRetries: 3,
	}

	var wg sync.WaitGroup
	wg.Add(1)

	go worker.Run(&wg)
	wg.Wait()
	failing := worker.Downloader.(*failingDownloader)
	expectedRetries := 3
	if failing.failCount != expectedRetries {
		t.Errorf("expected %d retries, but got %d", expectedRetries, failing.failCount)
	}
}