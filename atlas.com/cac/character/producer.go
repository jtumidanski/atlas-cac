package character

import (
   "atlas-cac/kafka/producers"
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

func CloseRangeAttack(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attackedAndDamaged byte, display byte, direction byte, stance byte, speed byte, damage map[uint32][]uint32) {
   producer := producers.ProduceEvent(l, "TOPIC_CLOSE_RANGE_ATTACK_EVENT")
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
