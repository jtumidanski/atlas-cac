package character

import (
	"atlas-cac/kafka/producers"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type closeRangeAttackEvent struct {
	WorldId            byte                `json:"worldId"`
	ChannelId          byte                `json:"channelId"`
	MapId              uint32              `json:"mapId"`
	CharacterId        uint32              `json:"characterId"`
	SkillId            uint32              `json:"skillId"`
	SkillLevel         byte                `json:"skillLevel"`
	AttackedAndDamaged byte                `json:"attackedAndDamaged"`
	Display            byte                `json:"display"`
	Direction          byte                `json:"direction"`
	Stance             byte                `json:"stance"`
	Speed              byte                `json:"speed"`
	Damage             map[uint32][]uint32 `json:"damage"`
}

func CloseRangeAttack(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attackedAndDamaged byte, display byte, direction byte, stance byte, speed byte, damage map[uint32][]uint32) {
	producer := producers.ProduceEvent(l, span, "TOPIC_CLOSE_RANGE_ATTACK_EVENT")
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attackedAndDamaged byte, display byte, direction byte, stance byte, speed byte, damage map[uint32][]uint32) {
		e := &closeRangeAttackEvent{
			WorldId:            worldId,
			ChannelId:          channelId,
			MapId:              mapId,
			CharacterId:        characterId,
			SkillId:            skillId,
			SkillLevel:         skillLevel,
			AttackedAndDamaged: attackedAndDamaged,
			Display:            display,
			Direction:          direction,
			Stance:             stance,
			Speed:              speed,
			Damage:             damage,
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}

type rangeAttackEvent struct {
	WorldId            byte                `json:"worldId"`
	ChannelId          byte                `json:"channelId"`
	MapId              uint32              `json:"mapId"`
	CharacterId        uint32              `json:"characterId"`
	SkillId            uint32              `json:"skillId"`
	SkillLevel         byte                `json:"skillLevel"`
	Stance             byte                `json:"stance"`
	AttackedAndDamaged byte                `json:"attackedAndDamaged"`
	Projectile         uint32              `json:"projectile"`
	Damage             map[uint32][]uint32 `json:"damage"`
	Speed              byte                `json:"speed"`
	Direction          byte                `json:"direction"`
	Display            byte                `json:"display"`
}

func RangeAttack(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attackedAndDamaged byte, display byte, direction byte, stance byte, speed byte, projectile uint32, damage map[uint32][]uint32) {
	producer := producers.ProduceEvent(l, span, "TOPIC_RANGE_ATTACK_EVENT")
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attackedAndDamaged byte, display byte, direction byte, stance byte, speed byte, projectile uint32, damage map[uint32][]uint32) {
		e := &rangeAttackEvent{
			WorldId:            worldId,
			ChannelId:          channelId,
			MapId:              mapId,
			CharacterId:        characterId,
			SkillId:            skillId,
			SkillLevel:         skillLevel,
			AttackedAndDamaged: attackedAndDamaged,
			Display:            display,
			Direction:          direction,
			Stance:             stance,
			Speed:              speed,
			Projectile:         projectile,
			Damage:             damage,
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}

type magicAttackEvent struct {
	WorldId            byte                `json:"worldId"`
	ChannelId          byte                `json:"channelId"`
	MapId              uint32              `json:"mapId"`
	CharacterId        uint32              `json:"characterId"`
	SkillId            uint32              `json:"skillId"`
	SkillLevel         byte                `json:"skillLevel"`
	Stance             byte                `json:"stance"`
	AttackedAndDamaged byte                `json:"attackedAndDamaged"`
	Damage             map[uint32][]uint32 `json:"damage"`
	Speed              byte                `json:"speed"`
	Direction          byte                `json:"direction"`
	Display            byte                `json:"display"`
	Charge             int32               `json:"charge"`
}

func MagicAttack(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attackedAndDamaged byte, display byte, direction byte, stance byte, speed byte, charge int32, damage map[uint32][]uint32) {
	producer := producers.ProduceEvent(l, span, "TOPIC_MAGIC_ATTACK_EVENT")
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attackedAndDamaged byte, display byte, direction byte, stance byte, speed byte, charge int32, damage map[uint32][]uint32) {
		e := &magicAttackEvent{
			WorldId:            worldId,
			ChannelId:          channelId,
			MapId:              mapId,
			CharacterId:        characterId,
			SkillId:            skillId,
			SkillLevel:         skillLevel,
			AttackedAndDamaged: attackedAndDamaged,
			Display:            display,
			Direction:          direction,
			Stance:             stance,
			Speed:              speed,
			Charge:             charge,
			Damage:             damage,
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}

type adjustManaEvent struct {
	CharacterId uint32 `json:"characterId"`
	Amount      int16  `json:"amount"`
}

func AdjustMana(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, amount int16) {
	producer := producers.ProduceEvent(l, span, "TOPIC_ADJUST_MANA")
	return func(characterId uint32, amount int16) {
		e := &adjustManaEvent{
			CharacterId: characterId,
			Amount:      amount,
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}

type mpEaterEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	CharacterId uint32 `json:"characterId"`
	SkillId     uint32 `json:"skillId"`
}

func ShowMPEater(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32) {
	producer := producers.ProduceEvent(l, span, "TOPIC_CHARACTER_MP_EATER_EVENT")
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32) {
		e := &mpEaterEvent{
			WorldId:     worldId,
			ChannelId:   channelId,
			MapId:       mapId,
			CharacterId: characterId,
			SkillId:     skillId,
		}
		producer(producers.CreateKey(int(characterId)), e)
	}
}
