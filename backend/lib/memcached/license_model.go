package memcached

import (
	"encoding/json"
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

type License struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type Licenses struct {
	Items []License `json:"items"`
}

func (licenses *Licenses) Save(client *memcache.Client) error {
	data, err := json.Marshal(licenses)
	if err != nil {
		return fmt.Errorf("Failed to serialize licenses.")
	}

	if err := client.Set(&memcache.Item{Key: LICENSES_CACHE_KEY, Value: data, Expiration: LICENSES_CACHE_RETENTION_SECONDS}); err != nil {
		return fmt.Errorf("Failed to cache licenses.")
	}
	return nil
}

func GetLicenses(client *memcache.Client) (*Licenses, error) {
	item, err := client.Get(LICENSES_CACHE_KEY)
	if err != nil {
		return nil, err
	}

	licenses := new(Licenses)
	if err := json.Unmarshal(item.Value, &licenses); err != nil {
		return nil, err
	}

	return licenses, nil
}
