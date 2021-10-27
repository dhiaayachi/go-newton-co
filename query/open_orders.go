package query

import (
	"net/http"
	"strconv"
	"time"
)

type OpenOrders struct {
	Limit int
	Offset int
	Symbol string
	TimeInForce TimeInForceAllowedValue
}

type OpenOrdersResponse []struct {
	OrderID      string    `json:"order_id"`
	Symbol       string    `json:"symbol"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	DateCreated  time.Time `json:"date_created"`
	OrderType    string    `json:"order_type"`
	TimeInForce  string    `json:"time_in_force"`
	Side         string    `json:"side"`
	QuantityLeft float64   `json:"quantity_left"`
	ExpiryTime   time.Time `json:"expiry_time"`
}

const openOrdersPath = "/order/open"

func (oo OpenOrders) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (oo OpenOrders) GetMethod() string {
	return http.MethodGet
}

func (oo OpenOrders) GetPath() string {
	return openOrdersPath
}

func (oo OpenOrders) GetParameters() []Parameter {
	params := make([]Parameter, 0)

	if oo.Limit != ANY {
		params = append(params, Parameter{string(Limit), strconv.Itoa(oo.Limit)})
	}

	if oo.Offset != ANY {
		params = append(params, Parameter{string(Offset), strconv.Itoa(oo.Offset)})
	}

	if oo.Symbol != NO_FILTER {
		params = append(params, Parameter{string(Symbol), oo.Symbol})
	}

	if string(oo.TimeInForce) != NO_FILTER {
		params = append(params, Parameter{string(TimeInForce), string(oo.TimeInForce)})
	}

	return params
}

func (oo OpenOrders) GetResponse() interface{} {
	return &OpenOrdersResponse{}
}

func (oo OpenOrders) IsPublic() bool {
	return false
}
