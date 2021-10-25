package query

type MinimumTradeAmounts struct {
}

func (mta MinimumTradeAmounts) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (mta MinimumTradeAmounts) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MinimumTradeAmounts) IsPublic() bool {
	return true
}
