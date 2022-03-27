package monster

import (
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

func GetById(l logrus.FieldLogger, span opentracing.Span) func(id uint32) (*Monster, error) {
	return func(id uint32) (*Monster, error) {
		resp, err := getById(id)(l, span)
		if err != nil {
			return nil, err
		}

		d := resp.Data()
		n := makeMonster(id, d.Attributes)
		return &n, nil
	}
}

func Damage(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, characterId uint32, damage uint32) {
	return emitDamage(l, span)
}

func makeMonster(id uint32, att attributes) Monster {
	return NewMonster(id, att.ControlCharacterId, att.MonsterId, att.X, att.Y, att.Stance, att.FH, att.Team, att.MaxHp, att.Hp, att.MaxMp, att.Mp)
}
