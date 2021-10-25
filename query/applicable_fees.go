package query

type ApplicableFees struct {
}

func (af ApplicableFees) GetBody() interface{} {
	return nil
}

func (af ApplicableFees) GetParameters() []Parameter {
	return []Parameter{}
}

func (af ApplicableFees) IsPublic() bool {
	return true
}
