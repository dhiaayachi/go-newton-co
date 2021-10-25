package query

type TickSizes struct {
}

func (ts TickSizes) GetParameters() []Parameter {
	return []Parameter{}
}

func (ts TickSizes) IsPublic() bool {
	return true
}
