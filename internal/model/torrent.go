package model

type Torrent struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	MagnetLink  string        `json:"magnetLink"`
	DownloadDir string        `json:"downloadDir"`
	Status      TorrentStatus `json:"status"`
}

type TorrentStatus string

const (
	StatusDownloading TorrentStatus = "downloading"
	StatusSeeding     TorrentStatus = "seeding"
	StatusCompleted   TorrentStatus = "completed"
	StatusError       TorrentStatus = "error"
)
