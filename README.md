# Scraper Bot

スクレイピング対象サイトから画像を収集し、ダウンロード後、Discordに通知するBotです。

## Features

- 複数Workerによる並列ダウンロード
- Chain of Responsibilityパターン採用
- リトライ戦略（Fixed / NoRetry / Exponential Backoff）を差し替え可能
- 簡単な環境変数設定だけで起動可能

## Setup

1. `.env`ファイルをプロジェクトルートに作成します（サンプルあり）
2. 必要なGoバージョンをインストールします（例：Go 1.21以上）
3. 以下のコマンドでビルド・実行できます

```bash
make build
make run
```

ディレクトリ構成
```
.
├── Makefile                  # ビルド・実行・クリーン用のコマンド定義
├── README.md                  # このプロジェクトの説明
├── cmd/
│   └── scraper/               # アプリケーションのエントリーポイント（main.go）
├── config/
│   └── config.go              # 環境変数読み込み・設定管理
├── downloads/                 # ダウンロードされた画像ファイルの保存先
├── go.mod, go.sum             # Go Modules管理ファイル
├── internal/
│   ├── app/
│   │   ├── downloader_app.go   # DownloaderAppの起動ロジック
│   │   └── worker/             # WorkerとHandlerの実装（Chain of Responsibilityパターン）
│   ├── di/
│   │   └── di.go               # 依存性注入（DI）管理
│   ├── domain/
│   │   ├── downloader/         # ダウンロードのインターフェース定義
│   │   ├── fetcher/            # スクレイピングのインターフェース定義
│   │   ├── logger/             # ロガーのインターフェース定義
│   │   └── uploader/           # Discordアップロード用インターフェース定義
│   └── service/
│       ├── downloader/         # ダウンロード処理の具体実装
│       ├── fetcher/            # スクレイピング処理の具体実装
│       ├── logger/             # ログ出力処理の具体実装
│       └── uploader/           # Discordへのアップロード処理の具体実装
└── scraper-bot                 # ビルド後のバイナリファイル