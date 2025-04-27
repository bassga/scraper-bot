package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

// FetchImageURLs は指定したURLにHTTPアクセスして、HTMLから画像のURL一覧を取得する
func FetchImageURLs(targetURL string) ([]string, error) {
	// URLにHTTPリクエストを送る
	resp, err := http.Get(targetURL)
	if err != nil {
		return nil, fmt.Errorf("failed to Get target URL: %w", err)
	}
	defer resp.Body.Close()

	// ステータスコードが200以外ならエラー
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// レスポンスボディから画像URLリストをパース
	return parseImageURLs(resp.Body, targetURL)
}

// parseImageURLs はHTMLの中からすべての<img>タグのsrc属性を抜き出してリスト化する
func parseImageURLs(r io.Reader, baseURL string) ([]string, error) {
	var urls []string

	doc, err := html.Parse(r)
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %w", err)
	}

	// HTMLツリーを深さ優先探索して画像URLを集める
	var traverse func(*html.Node, bool)
	traverse = func(n *html.Node, insideIE5 bool) {
		// divタグかチェック
		if n.Type == html.ElementNode && n.Data == "div" {
			for _, attr := range n.Attr {
				if attr.Key == "class" {
					if attr.Val == "rgn-container" {
						// rgn-containerに入ったら無視（探索しない）
						return
					} else if attr.Val == "ie5" {
						// ie5に入ったら以降の探索を有効にする
						insideIE5 = true
					}
				}
			}
		}

		// ie5内ならimgを収集
		if insideIE5 && n.Type == html.ElementNode && n.Data == "img" {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					absURL, err := toAbsoluteURL(baseURL, attr.Val)
					if err != nil {
						continue
					}
					urls = append(urls, absURL)
				}
			}
		}

		if n.Type == html.ElementNode && n.Data == "img" {
			for _, attr := range n.Attr {
				if attr.Key == "src" {
					absURL, err := toAbsoluteURL(baseURL, attr.Val)
					if err != nil {
						continue
					}
					urls = append(urls, absURL)
				}
			}
		}

		// 子ノードを再帰的に探索
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c, insideIE5)
		}
	}
	// ルートノードから探索開始
	traverse(doc, false)

	return urls, nil
}

func toAbsoluteURL(baseURL, imageURL string) (string, error) {
	parsedBase, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}
	parsedImage, err := url.Parse(imageURL)
	if err != nil {
		return "", err
	}
	absoluteURL := parsedBase.ResolveReference(parsedImage)
	return absoluteURL.String(), nil
}