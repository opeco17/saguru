package update

import "github.com/spf13/pflag"

type Options struct {
	Repository bool
	Issue      bool
	Cache      bool
	Index      bool
}

func GetOptions(flagSet *pflag.FlagSet) (*Options, error) {
	repository, err := flagSet.GetBool("repository")
	if err != nil {
		return nil, err
	}
	issue, err := flagSet.GetBool("issue")
	if err != nil {
		return nil, err
	}
	cache, err := flagSet.GetBool("cache")
	if err != nil {
		return nil, err
	}
	index, err := flagSet.GetBool("index")
	if err != nil {
		return nil, err
	}
	return &Options{Repository: repository, Issue: issue, Cache: cache, Index: index}, nil
}
