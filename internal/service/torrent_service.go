package service

import (
	"context"

	"github.com/nikitastetskiy/torrent-api/internal/model"
)

type TorrentService interface {
	Search(ctx context.Context, query string) ([]model.Torrent, error)
	Download(ctx context.Context, t model.Torrent) error
	Status(ctx context.Context, id string) (model.TorrentStatus, error)
	List(ctx context.Context) ([]model.Torrent, error)
}
