package query

import (
	"encoding/json"
	"net/http"
)

type CancelOrderBody struct {
	OrderId      string    `json:"order_id"`
}

type CancelOrder struct {
	Body CancelOrderBody
}

type CancelOrderResponse struct {
	OrderID      string    `json:"order_id"`
}

const cancelOrderPath = "/order/cancel"

func (co CancelOrder) GetBody() ([]byte, error) {
	body, err := json.Marshal(co.Body)
	if err != nil {
		return []byte(EMPTY_BODY), err
	}
	return body, nil
}

func (co CancelOrder) GetMethod() string {
	return http.MethodPost
}

func (co CancelOrder) GetPath() string {
	return cancelOrderPath
}

func (co CancelOrder) GetParameters() []Parameter {
	return []Parameter{}
}

func (co CancelOrder) GetResponse() interface{} {
	return &CancelOrderResponse{}
}

func (co CancelOrder) IsPublic() bool {
	return false
}
