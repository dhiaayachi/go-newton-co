package query

import "encoding/json"

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

func (no NewOrder) GetBody() ([]byte, error) {
	body, err := json.Marshal(no.Body)
	if err != nil {
		return []byte(EMPTY_BODY), err
	}
	return body, nil
}

func (no NewOrder) GetParameters() []Parameter {
	return []Parameter{}
}

func (no NewOrder) IsPublic() bool {
	return false
}
