package monster

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetMonster(l logrus.FieldLogger, span opentracing.Span) func(id uint32) (*Monster, error) {
	return func(id uint32) (*Monster, error) {
		resp, err := getById(l, span)(id)
		if err != nil {
			return nil, err
		}

		d := resp.Data()
		n := makeMonster(id, d.Attributes)
		return &n, nil
	}
}

func makeMonster(id uint32, att Attributes) Monster {
	return NewMonster(id, att.ControlCharacterId, att.MonsterId, att.X, att.Y, att.Stance, att.FH, att.Team, att.MaxHp, att.Hp, att.MaxMp, att.Mp)
}
