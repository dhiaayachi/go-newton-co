package query

type HealthCheck struct {
}

func (hc HealthCheck) GetBody() (string, error) {
	return EMPTY_BODY, nil
}

func (hc HealthCheck) GetParameters() []Parameter {
	return []Parameter{}
}

func (hc HealthCheck) IsPublic() bool {
	return true
}
