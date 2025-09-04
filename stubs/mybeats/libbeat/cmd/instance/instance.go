package instance

import "github.com/elastic/beats/v7/libbeat/publisher/processing"

type Settings struct {
	Name       string
	Version    string
	Processing processing.Supporter
}
