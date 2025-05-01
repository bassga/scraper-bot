package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/bassga/scraper-bot/config"
	"github.com/bassga/scraper-bot/internal/app"
	"github.com/bassga/scraper-bot/internal/di"
)


func main() {
	cfg := config.LoadConfig()
	container := di.NewContainer()
	downloaderApp := app.NewDownloaderApp(container, cfg.WorkerCount, cfg.MaxRetries)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	today := time.Now().Format("2006-01-02")
	downloadFolder := filepath.Join("downloads", today)
	err := os.MkdirAll(downloadFolder, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create downloads folder: %v", err)
		return
	}
	if err := downloaderApp.Run(ctx, cfg.TargetURL, downloadFolder); err != nil {
		log.Fatalf("application failed: %v", err)
	}
}