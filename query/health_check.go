package query

type HealthCheck struct {
}

func (hc HealthCheck) GetBody() ([]byte, error) {
	return []byte(EMPTY_BODY), nil
}

func (hc HealthCheck) GetParameters() []Parameter {
	return []Parameter{}
}

func (hc HealthCheck) IsPublic() bool {
	return true
}
