package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

// DownloadImage は指定した画像URLをダウンロードして保存する（タイムアウト付き）
func DownloadImage(url string, destFolder string, saveAsName string) (string, error) {
	// タイムアウト10秒のHTTPクライアント作成
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// HTTPリクエストを送る
	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	// ステータスコードチェック
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code while downloading image: %d", resp.StatusCode)
	}

	// 保存先ファイルパスを作成
	filePath := filepath.Join(destFolder, saveAsName)

	// ファイル作成
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// レスポンスボディ（画像データ）をファイルに書き込み
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save image to file: %w", err)
	}

	return filePath, nil
}