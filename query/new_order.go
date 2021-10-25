package query

import (
	"encoding/json"
	"net/http"
)

type NewOrderBody struct {
	OrderType   string  `json:"order_type"`
	TimeInForce TimeInForceAllowedValue  `json:"time_in_force"`
	Side        string  `json:"side"`
	Symbol      string  `json:"symbol"`
	Price       float64 `json:"price"`
	Quantity    float64 `json:"quantity"`
}

type NewOrder struct {
	Body NewOrderBody
}

const newOrderPath = "/order/new"

func (no NewOrder) GetBody() ([]byte, error) {
	body, err := json.Marshal(no.Body)
	if err != nil {
		return []byte(EMPTY_BODY), err
	}
	return body, nil
}

func (no NewOrder) GetMethod() string {
	return http.MethodPost
}

func (no NewOrder) GetPath() string {
	return newOrderPath
}

func (no NewOrder) GetParameters() []Parameter {
	return []Parameter{}
}

func (no NewOrder) IsPublic() bool {
	return false
}
