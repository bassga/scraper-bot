package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/bassga/scraper-bot/config"
	"github.com/bassga/scraper-bot/internal/downloader"
	"github.com/bassga/scraper-bot/internal/fetcher"
)

const workerCount = 5 // 同時に動かすワーカー（ダウンロード並列数）

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

	// ダウンロード対象を詰めるチャネル（キュー）を作成
	channels := make(chan struct {
		url string
		saveAsName string
	}, len(imageURLs)) // バッファあり（先に全部詰めるため）

	var wg sync.WaitGroup // 全ワーカーが終わるまで待つための仕組み

	// workerCount分のワーカーを起動
	for w := 0; w < workerCount; w++ {
		wg.Add(1)
		go func() { // ゴルーチン（軽量スレッド）で実行
			defer wg.Done() // このワーカーが終わったらカウントを減らす
			for channel := range channels { // チャネルから仕事を取り出す
				_, err := downloader.DownloadImage(channel.url, downloadFolder, channel.saveAsName)
				if err != nil {
					log.Printf("failed to download image: %v", err)
				} else {
					log.Printf("successfully downloaded: %s\n", channel.saveAsName)
				}
			}
		}()
	}

	// 仕事（画像URL＋ファイル名）をチャネルに流し込む
	for i, imageURL := range imageURLs {
		saveAsName := fmt.Sprintf("stamp_%03d.png", i+1) // stamp_001.png みたいな名前にする
		channels <- struct { url, saveAsName string } { 
			url: imageURL, saveAsName: saveAsName,
		}
	}

	close(channels) // 全部流し終わったのでチャネルを閉じる（これがないとワーカーが無限待ち状態になる）

	wg.Wait() // 全ワーカーが仕事終わるまで待つ

	log.Println("all images processed.") // 完了ログ
}