package main

import (
	"log"
	"os"

	"github.com/bassga/scraper-bot/config"
	"github.com/bassga/scraper-bot/internal/downloader"
	"github.com/bassga/scraper-bot/internal/fetcher"
)


func main() {
	// 設定情報読み込み
	cfg := config.LoadConfig()

	// 一時保存用フォルダを作成（なければ）
	const downloadFolder = "downloads"
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

	log.Printf("found %d images, starting download and upload...\n", len(imageURLs))

	// 各画像をダウンロード→Discordにアップロード
	for _, imageURL := range imageURLs {
		// 画像をダウンロード
		filePath, err := downloader.DownloadImage(imageURL, downloadFolder)
		if err != nil {
			log.Printf("failed to download image: %v", err)
			continue
		}
		// defer os.Remove(filePath) // 最後にダウンロードファイルを消す

		// // 画像をDiscordにアップロード
		// err = uploader.UploadImage(cfg.WebhookURL, filePath)
		// if err != nil {
		// 	log.Printf("failed to upload image: %v", err)
		// 	continue
		// }

		log.Printf("successfully uploaded: %s\n", filePath)
	}

	log.Println("all images processed.")
}