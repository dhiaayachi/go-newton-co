package query

type Symbols struct {
	BaseAsset, QuoteAsset string
}

func (s Symbols) GetBody() (string, error) {
	return EMPTY_BODY, nil
}

func (s Symbols) GetParameters() []Parameter {
	params := make([]Parameter, 0)

	if s.BaseAsset != NO_FILTER {
		params = append(params, Parameter{string(BaseAsset), s.BaseAsset})
	}

	if s.QuoteAsset != NO_FILTER {
		params = append(params, Parameter{string(QuoteAsset), s.QuoteAsset})
	}

	return params
}

func (s Symbols) IsPublic() bool {
	return true
}
