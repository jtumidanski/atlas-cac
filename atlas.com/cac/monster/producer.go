package monster

import (
	"atlas-cac/kafka/producers"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

type damageEvent struct {
	WorldId     byte   `json:"worldId"`
	ChannelId   byte   `json:"channelId"`
	MapId       uint32 `json:"mapId"`
	UniqueId    uint32 `json:"uniqueId"`
	CharacterId uint32 `json:"characterId"`
	Damage      uint32 `json:"damage"`
}

func Damage(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, characterId uint32, damage uint32) {
	producer := producers.ProduceEvent(l, span, "TOPIC_MONSTER_DAMAGE")
	return func(worldId byte, channelId byte, mapId uint32, uniqueId uint32, characterId uint32, damage uint32) {
		e := &damageEvent{
			WorldId:     worldId,
			ChannelId:   channelId,
			MapId:       mapId,
			UniqueId:    uniqueId,
			CharacterId: characterId,
			Damage:      damage,
		}
		producer(producers.CreateKey(int(uniqueId)), e)
	}
}
