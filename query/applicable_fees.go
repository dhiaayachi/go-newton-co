package query

type ApplicableFees struct {
}

func (af ApplicableFees) GetBody() (string, error) {
	return EMPTY_BODY, nil
}

func (af ApplicableFees) GetParameters() []Parameter {
	return []Parameter{}
}

func (af ApplicableFees) IsPublic() bool {
	return true
}
