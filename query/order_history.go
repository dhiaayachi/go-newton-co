package query

import "strconv"

type OrderHistory struct {
	Limit int
	Offset int
	StartDate int64
	EndDate int64
	Symbol string
	TimeInForce TimeInForceAllowedValue
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

func (oh OrderHistory) IsPublic() bool {
	return false
}
