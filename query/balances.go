package query

type Balances struct {
	Asset string
}

func (b Balances) GetBody() (string, error) {
	return EMPTY_BODY, nil
}

func (b Balances) GetParameters() []Parameter {
	params := make([]Parameter, 0)

	if b.Asset != NO_FILTER {
		params = append(params, Parameter{string(Asset), b.Asset})
	}

	return params
}

func (b Balances) IsPublic() bool {
	return false
}
