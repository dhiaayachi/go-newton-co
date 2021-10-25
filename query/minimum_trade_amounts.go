package query

type MinimumTradeAmounts struct {
}

func (mta MinimumTradeAmounts) GetBody() interface{} {
	return nil
}

func (mta MinimumTradeAmounts) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MinimumTradeAmounts) IsPublic() bool {
	return true
}
