package main

import (
	"log"

	"github.com/bassga/scraper-bot/config"
	"github.com/bassga/scraper-bot/internal/app"
	"github.com/bassga/scraper-bot/internal/di"
)

const workerCount = 5 // 同時に動かすワーカー（ダウンロード並列数）
const maxRetries = 3  // 最大リトライ回数

func main() {
	// 設定情報読み込み
	cfg := config.LoadConfig()

	container := di.NewContainer()

	downloaderApp := app.NewDownloaderApp(container, workerCount, maxRetries)

	if err := downloaderApp.Run(cfg.TargetURL); err != nil {
		log.Fatalf("application failed: %v", err)
	}
}