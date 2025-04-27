package uploader

// UploaderはローカルファイルをDiscordなどにアップロードする責務を持つインターフェース
type Uploader interface {
	UploadImage(webhookURL string, filePath string) error
}