package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/nikitastetskiy/torrent-api/internal/service"
)

func RegisterRoutes(r chi.Router) {
	r.Route("/torrents", func(r chi.Router) {
		r.Get("/", ListTorrents)
		r.Post("/search", SearchTorrents)
		r.Post("/download", DownloadTorrent)
		r.Get("/{id}/status", GetTorrentStatus)
	})
}

func ListTorrents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("[]"))
}

func SearchTorrents(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}

	results, err := service.SearchWolfmax4K(query)
	if err != nil {
		http.Error(w, "Search failed: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func DownloadTorrent(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Download a torrent"))
}

func GetTorrentStatus(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get torrent status"))
}
