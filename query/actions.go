package query

import "strconv"

type Actions struct {
	ActionType ActionType
	Limit int
	Offset int
	StartDate int64
	EndDate int64
}

func (arp Actions) GetParameters() []Parameter {
	params := make([]Parameter, 0)

	if arp.ActionType != ActionType(NO_FILTER) {
		params = append(params, Parameter{string(ActionTypeKey), string(arp.ActionType)})
	}

	if arp.Limit != int(ANY) {
		params = append(params, Parameter{string(Limit), strconv.Itoa(arp.Limit)})
	}

	if arp.Offset != int(ANY) {
		params = append(params, Parameter{string(Offset), strconv.Itoa(arp.Offset)})
	}

	if arp.StartDate != int64(ANY) {
		params = append(params, Parameter{string(StartDate), strconv.FormatInt(arp.StartDate, 10)})
	}

	if arp.EndDate != int64(ANY) {
		params = append(params, Parameter{string(EndDate), strconv.FormatInt(arp.EndDate, 10)})
	}

	return params
}

func (arp Actions) IsPublic() bool {
	return false
}
