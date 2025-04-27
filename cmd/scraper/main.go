package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/bassga/scraper-bot/config"
	"github.com/bassga/scraper-bot/internal/downloader"
	"github.com/bassga/scraper-bot/internal/fetcher"
)

func main() {
	// 設定情報読み込み
	cfg := config.LoadConfig()

	// 日付ごとに保存フォルダを分ける
	today := time.Now().Format("2006-01-02")
	downloadFolder := filepath.Join("downloads", today)

	// 保存フォルダを作成（なければ）
	err := os.MkdirAll(downloadFolder, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create downloads folder: %v", err)
	}

	// 対象ページから画像URL一覧を取得
	imageURLs, err := fetcher.FetchImageURLs(cfg.TargetURL)
	if err != nil {
		log.Fatalf("failed to fetch image URLs: %v", err)
	}

	// 画像URLが0件だったら終了
	if len(imageURLs) == 0 {
		log.Println("no images found on the target page.")
		return
	}

	log.Printf("found %d images, starting download...\n", len(imageURLs))

	// 各画像をダウンロード
	for i, imageURL := range imageURLs {
		saveAsName := fmt.Sprintf("stamp_%03d.png", i + 1)
		// 画像をダウンロード
		filePath, err := downloader.DownloadImage(imageURL, downloadFolder, saveAsName)
		if err != nil {
			log.Printf("failed to download image: %v", err)
			continue
		}

		log.Printf("successfully downloaded: %s\n", filePath)
	}

	log.Println("all images processed.")
}