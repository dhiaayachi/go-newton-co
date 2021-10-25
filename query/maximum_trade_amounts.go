package query

type MaximumTradeAmounts struct {
}

func (mta MaximumTradeAmounts) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (mta MaximumTradeAmounts) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MaximumTradeAmounts) IsPublic() bool {
	return true
}
