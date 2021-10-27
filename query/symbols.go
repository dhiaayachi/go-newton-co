package query

import "net/http"

type Symbols struct {
	BaseAsset, QuoteAsset string
}

type SymbolsResponse []string

const symbolsPath = "/symbols"

func (s Symbols) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (s Symbols) GetMethod() string {
	return http.MethodGet
}

func (s Symbols) GetPath() string {
	return symbolsPath
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

func (s Symbols) GetResponse() interface{} {
	return &SymbolsResponse{}
}

func (s Symbols) IsPublic() bool {
	return true
}
