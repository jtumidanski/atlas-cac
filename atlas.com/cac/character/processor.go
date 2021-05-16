package character

import (
   "atlas-cac/monster"
   "github.com/sirupsen/logrus"
)

func ProcessAttack(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attacked byte, damaged byte, attackedAndDamaged byte, stance byte, direction byte, rangedDirection byte, charge uint32, display byte, ranged bool, magic bool, speed byte, allDamage map[uint32][]uint32, x int16, y int16) error {
   return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attacked byte, damaged byte, attackedAndDamaged byte, stance byte, direction byte, rangedDirection byte, charge uint32, display byte, ranged bool, magic bool, speed byte, allDamage map[uint32][]uint32, x int16, y int16) error {
      if !ranged && !magic {
         CloseRangeAttack(l)(worldId, channelId, mapId, characterId, skillId, skillLevel, attackedAndDamaged, display, direction, stance, speed, allDamage)
      } else if ranged {
         // ranged
      } else if magic {
         // magic
      }

      //attackCount := uint32(1)

      for k, v := range allDamage {
         m, err := monster.GetMonster(k)
         if err != nil {
            l.WithError(err).Errorf("Cannot locate monster %d which the attack from %d hit.", k, characterId)
            continue
         }
         totalDamage := uint32(0)
         for _, e := range v {
            totalDamage += e
         }
         monster.Damage(l)(worldId, channelId, mapId, m.UniqueId(), characterId, totalDamage)
      }
      return nil
   }
}
