package query

import "net/http"

type HealthCheck struct {}

const healthCheckPath = "/health-check"

func (hc HealthCheck) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (hc HealthCheck) GetMethod() string {
	return http.MethodGet
}

func (hc HealthCheck) GetPath() string {
	return healthCheckPath
}

func (hc HealthCheck) GetParameters() []Parameter {
	return []Parameter{}
}

func (hc HealthCheck) IsPublic() bool {
	return true
}
