package worker

import (
	"context"
	"fmt"

	"github.com/bassga/scraper-bot/internal/domain/logger"
)

type ResizeHandler struct {
	base *BaseHandler
	Logger logger.Logger
	Folder string
}

func NewResizeHandler(logger logger.Logger, folder string) *ResizeHandler {
	return &ResizeHandler{
		base: &BaseHandler{},
		Logger: logger,
		Folder: folder,
	}
}

func (h *ResizeHandler) Handle(ctx context.Context, job Job) error {
	h.Logger.Info("resizeing image: %s", job.SaveAsName)
	// ここではダミーで成功とする
	fmt.Printf("[ResizeHandler] resized image: %s\n", job.SaveAsName)

	return h.base.Next(ctx, job)
}

func (h *ResizeHandler) SetNext(next JobHandler) {
	h.base.SetNext(next)
}
