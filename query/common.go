package query

type ActionType string

const (
	DEPOSIT    ActionType = "DEPOSIT"
	WITHDRAWAL ActionType = "WITHDRAWAL"
	TRANSACT   ActionType = "TRANSACT"
)

type QueryParameterKey string

const (
	Asset QueryParameterKey = "asset"
	ActionTypeKey QueryParameterKey = "action_type"
	Limit QueryParameterKey = "limit"
	Offset QueryParameterKey = "offset"
	StartDate QueryParameterKey = "start_date"
	EndDate QueryParameterKey = "end_date"
	Symbol QueryParameterKey = "symbol"
	TimeInForce QueryParameterKey = "time_in_force"
	BaseAsset = "base_asset"
	QuoteAsset = "quote_asset"
)

type TimeInForceAllowedValue string

const (
	NO_FILTER_VALUE TimeInForceAllowedValue = ""
	IOC TimeInForceAllowedValue = "IOC"
	GTC TimeInForceAllowedValue = "GTC"
	GTD TimeInForceAllowedValue = "GTD"
)

const (
	ANY int = -1
	NO_FILTER string = ""
	EMPTY_BODY = ""
)

type Parameter struct {
	Key   string
	Value string
}

type Query interface {
	GetBody() ([]byte, error)
	GetMethod() string
	GetPath() string
	GetParameters() []Parameter
	GetResponse() interface{}
	IsPublic() bool
}
