package character

import (
	"atlas-cac/kafka"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	consumerName = "character_attack_command"
	topicToken   = "TOPIC_CHARACTER_ATTACK_COMMAND"
)

type attackCommand struct {
	WorldId                  byte                `json:"worldId"`
	ChannelId                byte                `json:"channelId"`
	MapId                    uint32              `json:"mapId"`
	CharacterId              uint32              `json:"characterId"`
	NumberAttacked           byte                `json:"numberAttacked"`
	NumberDamaged            byte                `json:"numberDamaged"`
	NumberAttackedAndDamaged byte                `json:"NumberAttackedAndDamaged"`
	SkillId                  uint32              `json:"skillId"`
	SkillLevel               byte                `json:"skillLevel"`
	Stance                   byte                `json:"stance"`
	Direction                byte                `json:"direction"`
	RangedDirection          byte                `json:"rangedDirection"`
	Charge                   uint32              `json:"charge"`
	Display                  byte                `json:"display"`
	Ranged                   bool                `json:"ranged"`
	Magic                    bool                `json:"magic"`
	Speed                    byte                `json:"speed"`
	AllDamage                map[uint32][]uint32 `json:"allDamage"`
	X                        int16               `json:"x"`
	Y                        int16               `json:"y"`
}

func NewConsumer(groupId string) kafka.ConsumerConfig {
	return kafka.NewConsumerConfig[attackCommand](consumerName, topicToken, groupId, HandleAttackCommand())
}

func HandleAttackCommand() kafka.HandlerFunc[attackCommand] {
	return func(l logrus.FieldLogger, span opentracing.Span, command attackCommand) {
		err := ProcessAttack(l, span)(command.WorldId, command.ChannelId, command.MapId, command.CharacterId, command.SkillId, command.SkillLevel, command.NumberAttacked, command.NumberDamaged,
			command.NumberAttackedAndDamaged, command.Stance, command.Direction, command.RangedDirection, command.Charge, command.Display,
			command.Ranged, command.Magic, command.Speed, command.AllDamage, command.X, command.Y)
		if err != nil {
			l.WithError(err).Errorf("Unable to process attack for character %d.", command.CharacterId)
		}
	}
}
