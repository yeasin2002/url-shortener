package pkg

import (
	"errors"
	"time"
)

type URL struct {
	ID          string    `json:"id"`
	OriginalURL string    `json:"original_url"`
	ShortURL    string    `json:"short_url"`
	CreatedAt   time.Time `json:"created_at"`
}

var urlDB = make(map[string]URL)

func CreateUrl(originalUrl string) string {
	data, err := GenerateHash(originalUrl)
	if err != nil {
		println(err)
	}
	urlDB[data] = URL{
		ID:          data,
		OriginalURL: originalUrl,
		ShortURL:    data,
		CreatedAt:   time.Now(),
	}
	return data
}

func GetOriginalUrl(id string) (URL, error) {
	url, err := urlDB[id]
	if !err {
		println(err)
		return URL{}, errors.New("URL Not Found")
	}
	return url, nil
}
