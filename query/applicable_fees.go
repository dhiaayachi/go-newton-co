package query

import "net/http"

type ApplicableFees struct {}

type ApplicableFeesResponse struct {
	Maker float64 `json:"maker"`
	Taker float64 `json:"taker"`
}

const applicableFeesPath = "/fees"

func (af ApplicableFees) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (af ApplicableFees) GetMethod() string {
	return http.MethodGet
}

func (af ApplicableFees) GetPath() string {
	return applicableFeesPath
}

func (af ApplicableFees) GetParameters() []Parameter {
	return []Parameter{}
}

func (af ApplicableFees) GetResponse() interface{} {
	return &ApplicableFeesResponse{}
}

func (af ApplicableFees) IsPublic() bool {
	return true
}
