package di

import (
	downloaderDomain "github.com/bassga/scraper-bot/internal/domain/downloader"
	fetcherDomain "github.com/bassga/scraper-bot/internal/domain/fetcher"
	loggerDomain "github.com/bassga/scraper-bot/internal/domain/logger"
	uploaderDomain "github.com/bassga/scraper-bot/internal/domain/uploader"

	downloaderService "github.com/bassga/scraper-bot/internal/service/downloader"
	fetcherService "github.com/bassga/scraper-bot/internal/service/fetcher"
	loggerService "github.com/bassga/scraper-bot/internal/service/logger"
	uplaoderService "github.com/bassga/scraper-bot/internal/service/uploader"
)

type Container struct {
	Logger loggerDomain.Logger
	Fetcher fetcherDomain.Fetcher
	Downloader downloaderDomain.Downloader
	Uplaoder uploaderDomain.Uploader
}

func NewContainer() *Container {
	return &Container{
		Logger: loggerService.NewLogger(),
		Fetcher: fetcherService.NewFetcher(),
		Downloader: downloaderService.NewDownloader(),
		Uplaoder: uplaoderService.NewUploader(),
	}
}