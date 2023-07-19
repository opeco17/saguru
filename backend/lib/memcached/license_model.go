package memcached

import (
	"encoding/json"
	errorsutil "opeco17/saguru/lib/errors"

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
		return errorsutil.Wrap(err, err.Error())
	}

	if err := client.Set(&memcache.Item{Key: LICENSES_CACHE_KEY, Value: data, Expiration: LICENSES_CACHE_RETENTION_SECONDS}); err != nil {
		return errorsutil.Wrap(err, err.Error())
	}
	return nil
}

func GetLicenses(client *memcache.Client) (*Licenses, error) {
	item, err := client.Get(LICENSES_CACHE_KEY)
	if err != nil {
		return nil, errorsutil.Wrap(err, err.Error())
	}

	licenses := new(Licenses)
	if err := json.Unmarshal(item.Value, &licenses); err != nil {
		return nil, errorsutil.Wrap(err, err.Error())
	}

	return licenses, nil
}
