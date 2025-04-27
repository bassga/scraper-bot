package config

import "os"



type Config struct {
	TargetURL string
	WebhookURL string
}

func LoadConfig() *Config {
	targetURL := os.Getenv("TARGET_URL")
	
	webhookURL := os.Getenv("DISCORD_WEBHOOK_URL")
	
	return &Config{
		TargetURL: targetURL,
		WebhookURL: webhookURL,
	}
}