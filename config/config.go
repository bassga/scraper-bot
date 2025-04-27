package config

import "os"



type Config struct {
	TargetURL string
	WebhookURL string
}

func LoadConfig() *Config {
	targetURL := os.Getenv("TARGET_URL")
	if targetURL == "" {
		targetURL = "https://gbf-wiki.com/index.php?%E3%82%B9%E3%82%BF%E3%83%B3%E3%83%97%E4%B8%80%E8%A6%A7"
	}

	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	if webhookURL == "" {
		webhookURL = "https://discord.com/api/webhooks/1365959000356814941/Z-3c8dw1m8jkJPgB0RcyELwbCax4bmC5F390jAKd5EagIUmyoFL1VM9zQ8psvf8Z19Hg"
	}

	return &Config{
		TargetURL: targetURL,
		WebhookURL: webhookURL,
	}
}