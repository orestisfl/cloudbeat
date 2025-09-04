package beat

import (
	"time"
	"github.com/elastic/elastic-agent-libs/mapstr"
	"errors"
	"github.com/elastic/beats/v7/libbeat/common"
	"github.com/elastic/beats/v7/libbeat/processors"
	"github.com/elastic/elastic-agent-libs/config"
)

var GracefulExit = errors.New("graceful exit")

type Beater interface {
	Run(*Beat) error
	Stop()
}

type Beat struct {
	Manager Manager
	Publisher Pipeline
	Registry *common.Registry
}

type Creator func(*Beat, *config.C) (Beater, error)

type Event struct {
	Meta      mapstr.M
	Timestamp time.Time
	Fields    mapstr.M
}

type Processor interface{}
type Manager interface {
	Start() error
	Enabled() bool
	Stop()
	UpdateStatus(status string, msg string)
}

func (e *Event) PutValue(key string, value interface{}) (mapstr.M, error) {
	e.Fields[key] = value
	return e.Fields, nil
}

type Client interface{
	PublishAll([]Event)
	Close() error
}

type Pipeline struct{}

func (p *Pipeline) ConnectWith(config ClientConfig) (Client, error) {
	return nil, nil
}

type ClientConfig struct{
	Processing ProcessingConfig
}

type ProcessingConfig struct{
	Processor *processors.Processors
}
