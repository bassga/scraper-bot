package app

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/bassga/scraper-bot/internal/app/worker"
	"github.com/bassga/scraper-bot/internal/di"
)

type DownloaderApp struct {
	container             *di.Container
	workerCount, maxRetries int
}

func NewDownloaderApp(container *di.Container, workerCount, maxRetries int) *DownloaderApp {
	return &DownloaderApp{
		container:   container,
		workerCount: workerCount,
		maxRetries:  maxRetries,
	}
}

func (app *DownloaderApp) Run(ctx context.Context, targetURL , downloadFolder string) error {
	
	err := os.MkdirAll(downloadFolder, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to create downloads folder: %w", err)
	}

	imageURLs, err := app.container.Fetcher.FetchImageURLs(targetURL)
	if err != nil {
		return fmt.Errorf("failed to fetch image URLs: %w", err)
	}

	if len(imageURLs) == 0 {
		app.container.Logger.Info("no images found on the target page.")
		return nil
	}

	app.container.Logger.Info("found %d images, starting download...\n", len(imageURLs))

	jobs := make(chan worker.Job, len(imageURLs))

	var wg sync.WaitGroup

	for w := 0; w < app.workerCount; w++ {
		wg.Add(1)

		// ハンドラチェーンの構築
		downloadHandler := worker.NewDownloadHandler(
			app.container.Downloader,
			app.container.Logger,
			downloadFolder,
		)

		resizeHandler := worker.NewResizeHandler(
			app.container.Logger,
			downloadFolder,
		)

		downloadHandler.SetNext(resizeHandler)

		worker := &worker.Worker{
			Downloader:    app.container.Downloader,
			Logger:        app.container.Logger,
			Ctx:           ctx,
			Jobs:          jobs,
			Folder:        downloadFolder,
			MaxRetries: 	 app.maxRetries,
			RetryStrategy: &worker.FixedRetryStrategy{},
			JobHandler:    downloadHandler, // ハンドラチェーンの開始点
		}

		go worker.Run(&wg)
	}

	for i, imageURL := range imageURLs {
		saveAsName := fmt.Sprintf("stamp_%03d.png", i+1)
		jobs <- worker.Job{URL: imageURL, SaveAsName: saveAsName}
	}

	close(jobs)
	wg.Wait()

	app.container.Logger.Info("all images processed.")
	return nil
}
