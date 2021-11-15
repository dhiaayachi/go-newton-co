package query

import (
	"net/http"
	"strconv"
	"time"
)

type OrderHistory struct {
	Limit int
	Offset int
	StartDate int64
	EndDate int64
	Symbol string
	TimeInForce TimeInForceAllowedValue
}

type OrderHistoryResponse []struct {
	OrderID      string    `json:"order_id"`
	Symbol       string    `json:"symbol"`
	Quantity     int       `json:"quantity"`
	Price        float64   `json:"price"`
	Status       string    `json:"status"`
	DateCreated  time.Time `json:"date_created"`
	DateExecuted string    `json:"date_executed"`
	OrderType    string    `json:"order_type"`
	TimeInForce  string    `json:"time_in_force"`
	Side         string    `json:"side"`
	ExpiryTime   time.Time `json:"expiry_time,omitempty"`
}

const orderHistoryPath = "/order/history"

func (oh OrderHistory) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (oh OrderHistory) GetMethod() string {
	return http.MethodGet
}

func (oh OrderHistory) GetPath() string {
	return orderHistoryPath
}

func (oh OrderHistory) GetParameters() []Parameter {
	params := make([]Parameter, 0)

	if oh.Limit != ANY {
		params = append(params, Parameter{string(Limit), strconv.Itoa(oh.Limit)})
	}

	if oh.Offset != ANY {
		params = append(params, Parameter{string(Offset), strconv.Itoa(oh.Offset)})
	}

	if oh.StartDate != int64(ANY) {
		params = append(params, Parameter{string(StartDate), strconv.Itoa(int(oh.StartDate))})
	}

	if oh.EndDate != int64(ANY) {
		params = append(params, Parameter{string(EndDate), strconv.Itoa(int(oh.EndDate))})
	}

	if oh.Symbol != NO_FILTER {
		params = append(params, Parameter{string(Symbol), oh.Symbol})
	}

	if string(oh.TimeInForce) != NO_FILTER {
		params = append(params, Parameter{string(TimeInForce), string(oh.TimeInForce)})
	}

	return params
}

func (oh OrderHistory) GetResponse() interface{} {
	return &OrderHistoryResponse{}
}

func (oh OrderHistory) IsPublic() bool {
	return false
}
