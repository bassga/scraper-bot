package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	TargetURL  string
	WebhookURL string
}

func LoadConfig() *Config {
	// 最初に.envファイル読み込む
	err := godotenv.Load()
	if err != nil {
		log.Println(".envファイルが見つかりませんでした（環境変数だけ読み込みます）")
	}

	targetURL := os.Getenv("TARGET_URL")
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")

	return &Config{
		TargetURL:  targetURL,
		WebhookURL: webhookURL,
	}
}