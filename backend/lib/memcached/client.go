package memcached

import (
	"strconv"

	errorsutil "opeco17/saguru/lib/errors"

	"github.com/bradfitz/gomemcache/memcache"
)

func GetMemcachedClient(host string, port int) (*memcache.Client, error) {
	if host == "" {
		return nil, errorsutil.CustomError{Message: "You must set 'host' to connect to MongoDB"}
	}
	if port < 0 {
		return nil, errorsutil.CustomError{Message: "Port number is invalid"}
	}
	return memcache.New(host + ":" + strconv.Itoa(port)), nil
}
