package query

import "strconv"

type OpenOrders struct {
	Limit int
	Offset int
	Symbol string
	TimeInForce TimeInForceAllowedValue
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

func (oo OpenOrders) IsPublic() bool {
	return false
}
