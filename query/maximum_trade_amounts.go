package query

type MaximumTradeAmounts struct {
}

func (mta MaximumTradeAmounts) GetBody() (string, error) {
	return EMPTY_BODY, nil
}

func (mta MaximumTradeAmounts) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MaximumTradeAmounts) IsPublic() bool {
	return true
}
