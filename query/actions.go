package query

import (
	"net/http"
	"strconv"
	"time"
)

type Actions struct {
	ActionType ActionType
	Limit int
	Offset int
	StartDate int64
	EndDate int64
}

type ActionsResponse []struct {
	Type        string    `json:"type"`
	Asset       string    `json:"asset"`
	Quantity    float64   `json:"quantity"`
	Status      string    `json:"status"`
	DateCreated time.Time `json:"date_created"`
	Price       float64   `json:"price,omitempty"`
}

const actionsPath = "/actions"

func (a Actions) GetMethod() string {
	return http.MethodGet
}

func (a Actions) GetPath() string {
	return actionsPath
}

func (a Actions) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (a Actions) GetParameters() []Parameter {
	params := make([]Parameter, 0)

	if a.ActionType != ActionType(NO_FILTER) {
		params = append(params, Parameter{string(ActionTypeKey), string(a.ActionType)})
	}

	if a.Limit != int(ANY) {
		params = append(params, Parameter{string(Limit), strconv.Itoa(a.Limit)})
	}

	if a.Offset != int(ANY) {
		params = append(params, Parameter{string(Offset), strconv.Itoa(a.Offset)})
	}

	if a.StartDate != int64(ANY) {
		params = append(params, Parameter{string(StartDate), strconv.FormatInt(a.StartDate, 10)})
	}

	if a.EndDate != int64(ANY) {
		params = append(params, Parameter{string(EndDate), strconv.FormatInt(a.EndDate, 10)})
	}

	return params
}

func (a Actions) GetResponse() interface{} {
	return &ActionsResponse{}
}

func (a Actions) IsPublic() bool {
	return false
}
