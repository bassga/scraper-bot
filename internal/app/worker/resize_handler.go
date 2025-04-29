package worker

import (
	"context"
	"fmt"

	"github.com/bassga/scraper-bot/internal/domain/logger"
)


type ResizeHandler struct {
	BaseHandler
	Logger logger.Logger
	Folder string
}

func (h *ResizeHandler) Handle(ctx context.Context, job Job) error {
	h.Logger.Info("resizeing image: %s", job.SaveAsName)
	// ここではダミーで成功とする
	fmt.Printf("[ResizeHandler] resized image: %s\n", job.SaveAsName)

	return h.Next(ctx, job)
}

func (h *ResizeHandler) SetNext(next JobHandler) {
	h.BaseHandler.SetNext(next)
}
