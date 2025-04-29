package worker

import (
	"context"

	"github.com/bassga/scraper-bot/internal/domain/downloader"
	"github.com/bassga/scraper-bot/internal/domain/logger"
)


type DownloadHandler struct {
	BaseHandler
	Downloader downloader.Downloader
	Logger logger.Logger
	Folder string
}

func (h *DownloadHandler) Handle(ctx context.Context, job Job) error {
	_, err := h.Downloader.DownloadImage(ctx, job.URL, h.Folder, job.SaveAsName)
	if err != nil {
		h.Logger.Error("download failed: %v", err)
		return err
	}
	h.Logger.Info("download succeded: %s", job.SaveAsName)
	return h.Next(ctx, job)
}

func (h *DownloadHandler) SetNext(next JobHandler) {
	h.BaseHandler.SetNext(next)
}
