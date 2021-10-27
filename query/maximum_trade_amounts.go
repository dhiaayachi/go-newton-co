package query

import "net/http"

type MaximumTradeAmounts struct {}

type MaximumTradeAmountsResponse map[string]struct {
	Buy  float64 `json:"buy"`
	Sell float64 `json:"sell"`
}

const maximumTradeAmountsPath = "/order/maximums"

func (mta MaximumTradeAmounts) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (mta MaximumTradeAmounts) GetMethod() string {
	return http.MethodGet
}

func (mta MaximumTradeAmounts) GetPath() string {
	return maximumTradeAmountsPath
}

func (mta MaximumTradeAmounts) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MaximumTradeAmounts) GetResponse() interface{} {
	return &MaximumTradeAmountsResponse{}
}

func (mta MaximumTradeAmounts) IsPublic() bool {
	return true
}
