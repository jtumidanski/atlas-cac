package character

type Model struct {
	id    uint32
	jobId uint16
	hp    uint16
}

func (a Model) Id() uint32 {
	return a.id
}

func (a Model) HP() uint16 {
	return a.hp
}

func (a Model) JobId() uint16 {
	return a.jobId
}
