package common

import "github.com/elastic/elastic-agent-libs/config"

type MapStr map[string]interface{}

type Config = config.C

type Registry struct{}

func (r *Registry) MustRegisterInput(input interface{}) {}
