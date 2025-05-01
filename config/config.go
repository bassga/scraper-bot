package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	TargetURL  string
	WebhookURL string
	WorkerCount int
	MaxRetries int
}

func LoadConfig() *Config {
	// 最初に.envファイル読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println(".envファイルが見つかりませんでした（環境変数だけ読み込みます）")
	}

	targetURL := os.Getenv("TARGET_URL")
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	workerCount, err := strconv.Atoi(os.Getenv("WORKER_COUNT"))
	if err != nil {
		log.Println("WORKER_COUNTが不正です。デフォルト値3を使います")
		workerCount = 3
	}

	maxRetries, err := strconv.Atoi(os.Getenv("MAX_RETRIES"))
	if err != nil {
		log.Println("MAX_RETRIESが不正です。デフォルト値3を使います")
		maxRetries = 3
	}

	return &Config{
		TargetURL:   targetURL,
		WebhookURL:  webhookURL,
		WorkerCount: workerCount,
		MaxRetries:  maxRetries,
	}
}