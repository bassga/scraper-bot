package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/bassga/scraper-bot/internal/domain/downloader"
)

// DownloaderImpl は Downloader インターフェースの実装
type DownloaderImpl struct{}

// NewDownloader は DownloaderImpl のコンストラクタ
func NewDownloader() downloader.Downloader {
	return &DownloaderImpl{}
}

// DownloadImage は指定したURLの画像を保存先フォルダに保存する
func (d *DownloaderImpl) DownloadImage(url string, destFolder string, saveAsName string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code while downloading image: %d", resp.StatusCode)
	}

	filePath := filepath.Join(destFolder, saveAsName)
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save image to file: %w", err)
	}

	return filePath, nil
}