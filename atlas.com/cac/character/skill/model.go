package skill

type Model struct {
	id          uint32
	level       uint8
	masterLevel uint8
	expiration  int64
	hidden      bool
	fourthJob   bool
}

func NewModel(id uint32, level uint8, masterLevel uint8, expiration int64, hidden bool, fourthJob bool) Model {
	return Model{
		id:          id,
		level:       level,
		masterLevel: masterLevel,
		expiration:  expiration,
		hidden:      hidden,
		fourthJob:   fourthJob,
	}
}

func (s Model) Hidden() bool {
	return s.hidden
}

func (s Model) Id() uint32 {
	return s.id
}

func (s Model) Level() uint8 {
	return s.level
}

func (s Model) Expiration() int64 {
	return s.expiration
}

func (s Model) FourthJob() bool {
	return s.fourthJob
}

func (s Model) MasterLevel() uint8 {
	return s.masterLevel
}
