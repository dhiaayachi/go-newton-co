package query

type HealthCheck struct {
}

func (hc HealthCheck) GetParameters() []Parameter {
	return []Parameter{}
}

func (hc HealthCheck) IsPublic() bool {
	return true
}
