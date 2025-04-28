package service

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// TorrentResult represents one search result
type TorrentResult struct {
	Name       string `json:"name"`
	MagnetLink string `json:"magnetLink"`
	Size       string `json:"size,omitempty"`
}

// SearchWolfmax4K searches torrents in Wolfmax4K
func SearchWolfmax4K(query string) ([]TorrentResult, error) {
	searchQuery := strings.ReplaceAll(query, " ", "+")
	url := fmt.Sprintf("https://wolfmax4k.com/buscar/%s", searchQuery)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; torrent-manager/1.0)")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []TorrentResult

	// Adjust this to the real HTML structure
	doc.Find(".torrent-item").Each(func(i int, s *goquery.Selection) {
		name := s.Find(".torrent-title").Text()
		magnet, _ := s.Find(".magnet-link").Attr("href")
		size := s.Find(".torrent-size").Text()

		if magnet != "" {
			results = append(results, TorrentResult{
				Name:       strings.TrimSpace(name),
				MagnetLink: strings.TrimSpace(magnet),
				Size:       strings.TrimSpace(size),
			})
		}
	})

	return results, nil
}
