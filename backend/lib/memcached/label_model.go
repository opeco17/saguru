package memcached

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type Label struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Labels struct {
	Items []Label `json:"items"`
}

func (labels *Labels) Save(client *memcache.Client) error {
	data, err := json.Marshal(labels)
	if err != nil {
		return fmt.Errorf("Failed to serialize labels.")
	}

	if err := client.Set(&memcache.Item{Key: LABELS_CACHE_KEY, Value: data, Expiration: LABELS_CACHE_RETENTION_SECONDS}); err != nil {
		return fmt.Errorf("Failed to cache labels.")
	}
	return nil
}

func GetLabels(client *memcache.Client) (*Labels, error) {
	item, err := client.Get(LABELS_CACHE_KEY)
	if err != nil {
		return nil, err
	}

	labels := new(Labels)
	if err := json.Unmarshal(item.Value, &labels); err != nil {
		return nil, err
	}

	return labels, nil
}
