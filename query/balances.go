package query

import "net/http"

type Balances struct {
	Asset string
}

type BalancesResponse map[string]float64

const balancesPath = "/balances"

func (b Balances) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (b Balances) GetMethod() string {
	return http.MethodGet
}

func (b Balances) GetPath() string {
	return balancesPath
}

func (b Balances) GetParameters() []Parameter {
	params := make([]Parameter, 0)

	if b.Asset != NO_FILTER {
		params = append(params, Parameter{string(Asset), b.Asset})
	}

	return params
}

func (b Balances) GetResponse() interface{} {
	return &Balances{}
}

func (b Balances) IsPublic() bool {
	return false
}
