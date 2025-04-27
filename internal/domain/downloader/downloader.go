package downloader

// Downloaderは画像を指定先にダウンロードする責務を持つインターフェース
type Downloader interface {
	DownloadImage(url string, destFolder string, saveAsName string) (string, error)
}