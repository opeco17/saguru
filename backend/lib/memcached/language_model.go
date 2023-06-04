package memcached

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type Language struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Languages struct {
	Items []Language `json:"items"`
}

func (languages *Languages) Save(client *memcache.Client) error {
	data, err := json.Marshal(languages)
	if err != nil {
		return fmt.Errorf("Failed to serialize languages.")
	}

	if err := client.Set(&memcache.Item{Key: LANGUAGES_CACHE_KEY, Value: data, Expiration: LANGUAGES_CACHE_RETENTION_SECONDS}); err != nil {
		return fmt.Errorf("Failed to cache languages.")
	}
	return nil
}

func GetLanguages(client *memcache.Client) (*Languages, error) {
	item, err := client.Get(LANGUAGES_CACHE_KEY)
	if err != nil {
		return nil, err
	}

	languages := new(Languages)
	if err := json.Unmarshal(item.Value, &languages); err != nil {
		return nil, err
	}

	return languages, nil
}
