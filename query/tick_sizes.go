package query

import "net/http"

type TickSizes struct {}

const tickSizesPath = "/order/tick-sizes"

func (ts TickSizes) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (ts TickSizes) GetMethod() string {
	return http.MethodGet
}

func (ts TickSizes) GetPath() string {
	return tickSizesPath
}

func (ts TickSizes) GetParameters() []Parameter {
	return []Parameter{}
}

func (ts TickSizes) IsPublic() bool {
	return true
}
