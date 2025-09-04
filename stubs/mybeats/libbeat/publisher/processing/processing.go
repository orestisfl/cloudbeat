package processing

type Supporter interface{}

func MakeDefaultSupport(b bool, c chan struct{}, s ...Supporter) Supporter {
	return nil
}

func WithECS() Supporter {
	return nil
}

func WithAgentMeta() Supporter {
	return nil
}