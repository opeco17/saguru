package memcached

import (
	"fmt"
	"strconv"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/sirupsen/logrus"
)

func GetMemcachedClient(host string, port int) (*memcache.Client, error) {
	if host == "" {
		err_message := "You must set 'host' to connect to MongoDB"
		logrus.Error(err_message)
		return nil, fmt.Errorf(err_message)
	}
	if port < 0 {
		err_message := "Port number is invalid"
		logrus.Error(err_message)
		return nil, fmt.Errorf(err_message)
	}
	return memcache.New(host + ":" + strconv.Itoa(port)), nil
}
