package downloader

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// DownloadImage は指定した画像URLをダウンロードして保存する
func DownloadImage(url string, destFolder string, saveAsName string) (string, error) {
	// HTTPリクエストで画像データ取得
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to download image: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code while downloading image: %d", resp.StatusCode)
	}

	// ファイル名をURLから推測
	filePath := filepath.Join(destFolder, saveAsName)

	// 保存先ファイル作成
	out, err := os.Create(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// ダウンロードしたデータを書き込み
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to save image to file: %w", err)
	}

	return filePath, nil
}