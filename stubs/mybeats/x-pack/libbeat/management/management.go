package management

import (
	"github.com/elastic/beats/v7/libbeat/common/reload"
	"github.com/elastic/elastic-agent-client/v7/pkg/client"
	"github.com/elastic/elastic-agent-client/v7/pkg/proto"
)

var ConfigTransform = &configTransform{}

type configTransform struct{}

func (t *configTransform) SetTransform(f func(rawIn *proto.UnitExpectedConfig, agentInfo *client.AgentInfo) ([]*reload.ConfigWithMeta, error)) {
}

func CreateInputsFromStreams(rawIn *proto.UnitExpectedConfig, s string, agentInfo *client.AgentInfo) ([]map[string]interface{}, error) {
	return nil, nil
}

func CreateReloadConfigFromInputs(inputs []map[string]interface{}) ([]*reload.ConfigWithMeta, error) {
	return nil, nil
}
