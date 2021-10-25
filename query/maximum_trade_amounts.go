package query

type MaximumTradeAmounts struct {
}

func (mta MaximumTradeAmounts) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MaximumTradeAmounts) IsPublic() bool {
	return true
}
