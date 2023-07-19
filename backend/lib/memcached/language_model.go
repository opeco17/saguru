package memcached

import (
	"encoding/json"

	errorsutil "opeco17/saguru/lib/errors"

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
		return errorsutil.Wrap(err, err.Error())
	}

	if err := client.Set(&memcache.Item{Key: LANGUAGES_CACHE_KEY, Value: data, Expiration: LANGUAGES_CACHE_RETENTION_SECONDS}); err != nil {
		return errorsutil.Wrap(err, err.Error())
	}
	return nil
}

func GetLanguages(client *memcache.Client) (*Languages, error) {
	item, err := client.Get(LANGUAGES_CACHE_KEY)
	if err != nil {
		return nil, errorsutil.Wrap(err, err.Error())
	}

	languages := new(Languages)
	if err := json.Unmarshal(item.Value, &languages); err != nil {
		return nil, errorsutil.Wrap(err, err.Error())
	}

	return languages, nil
}
