package processors

func RegisterPlugin(name string, constructor interface{}) {}

type PluginConfig []interface{}

type Processors struct{}

func New(config PluginConfig) (*Processors, error) {
	return &Processors{}, nil
}