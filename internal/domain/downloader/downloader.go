package downloader

import "context"

// Downloaderは画像を指定先にダウンロードする責務を持つインターフェース
type Downloader interface {
	DownloadImage(ctx context.Context, url string, destFolder string, saveAsName string) (string, error)
}