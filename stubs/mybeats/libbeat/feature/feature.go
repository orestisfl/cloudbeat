package feature

import "github.com/elastic/beats/v7/libbeat/common"

type Factory func(string, *common.Config, string) (Reloadable, error)

type Reloadable interface{}

type Feature struct{}