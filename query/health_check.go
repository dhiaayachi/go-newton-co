package query

type HealthCheck struct {
}

func (hc HealthCheck) GetBody() interface{} {
	return nil
}

func (hc HealthCheck) GetParameters() []Parameter {
	return []Parameter{}
}

func (hc HealthCheck) IsPublic() bool {
	return true
}
