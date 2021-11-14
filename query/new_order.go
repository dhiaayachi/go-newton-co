package query

import (
	"encoding/json"
	"net/http"
	"time"
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

type NewOrderResponse struct {
	OrderId      string    `json:"order_id"`
	Symbol       string    `json:"symbol"`
	Quantity     float64       `json:"quantity"`
	Price        float64   `json:"price"`
	OrderType    string    `json:"order_type"`
	TimeInForce  string    `json:"time_in_force"`
	Side         string    `json:"side"`
	DateCreated  time.Time `json:"date_created"`
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

func (no NewOrder) GetResponse() interface{} {
	return &NewOrderResponse{}
}

func (no NewOrder) IsPublic() bool {
	return false
}
