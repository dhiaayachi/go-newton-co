package query

import "net/http"

type MinimumTradeAmounts struct {}

const minimumTradeAmountsPath = "/order/minimums"

func (mta MinimumTradeAmounts) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (mta MinimumTradeAmounts) GetMethod() string {
	return http.MethodGet
}

func (mta MinimumTradeAmounts) GetPath() string {
	return minimumTradeAmountsPath
}

func (mta MinimumTradeAmounts) GetParameters() []Parameter {
	return []Parameter{}
}

func (mta MinimumTradeAmounts) IsPublic() bool {
	return true
}
