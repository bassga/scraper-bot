package worker

import (
	"context"

	"github.com/bassga/scraper-bot/internal/domain/downloader"
	"github.com/bassga/scraper-bot/internal/domain/logger"
)


type DownloadHandler struct {
	base *BaseHandler
	Downloader downloader.Downloader
	Logger logger.Logger
	Folder string
}

func NewDownloadHandler(dl downloader.Downloader, logger logger.Logger, folder string) *DownloadHandler {
	return &DownloadHandler{
		base:       &BaseHandler{},
		Downloader: dl,
		Logger:     logger,
		Folder:     folder,
	}
}

func (h *DownloadHandler) Handle(ctx context.Context, job Job) error {
	_, err := h.Downloader.DownloadImage(ctx, job.URL, h.Folder, job.SaveAsName)
	if err != nil {
		h.Logger.Error("download failed: %v", err)
		return err
	}
	h.Logger.Info("download succeded: %s", job.SaveAsName)
	return h.base.Next(ctx, job)
}

func (h *DownloadHandler) SetNext(next JobHandler) {
	h.base.SetNext(next)
}
