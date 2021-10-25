package query

type TickSizes struct {
}

func (ts TickSizes) GetBody() (string, error) {
	return EMPTY_BODY, nil
}

func (ts TickSizes) GetParameters() []Parameter {
	return []Parameter{}
}

func (ts TickSizes) IsPublic() bool {
	return true
}
