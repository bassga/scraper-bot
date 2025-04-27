package app

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/bassga/scraper-bot/internal/di"
)


type DownloaderApp struct {
	container *di.Container
	workerCount, maxRetries int
}

func NewDownloaderApp(container *di.Container, workerCount, maxRetries int) *DownloaderApp {
	return &DownloaderApp{
		container: container,
		workerCount: workerCount,
		maxRetries: maxRetries,
	}
}

func (app *DownloaderApp) Run(targetURL string) error {
	today := time.Now().Format("2006-01-02")
	downloadFolder := filepath.Join("downloads", today)

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

	channels := make(chan struct {
		url string
		saveAsName string
	}, len(imageURLs))

	var wg sync.WaitGroup

	for w := 0; w < app.workerCount; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for channel := range channels {
				var filePath string
				var err error
				for attempt := 1; attempt <= app.maxRetries; attempt++ {
					filePath, err = app.container.Downloader.DownloadImage(channel.url, downloadFolder, channel.saveAsName)
					if err == nil {
						break
					}
					app.container.Logger.Error("failed to download image (attempt %d/%d): %v", attempt, app.maxRetries, err)
					time.Sleep(1 * time.Second)
				}

				if err != nil {
					app.container.Logger.Error("giving up downloading image after retries: %v", err)
					continue
				}

				app.container.Logger.Info("successfully downloaded: %s\n", filePath)
			}
		}()
	}

	for i, imageURL := range imageURLs {
		saveAsName := fmt.Sprintf("stamp_%03d.png", i+1)
		channels <- struct {
			url string
			saveAsName string
		}{url: imageURL, saveAsName: saveAsName}
	}

	close(channels)
	wg.Wait()

	app.container.Logger.Info("all images processed.")
	return nil
}