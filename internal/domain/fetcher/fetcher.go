package fetcher

// Fetcherは画像URLリストを取得する責務を持つインターフェース
type Fetcher interface {
	FetchImageURLs(targetURL string) ([]string, error)
}
