package query

type MinimumTradeAmount struct {
}

func (mta MinimumTradeAmount) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MinimumTradeAmount) IsPublic() bool {
	return true
}
