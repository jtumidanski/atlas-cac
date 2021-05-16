package character

import (
   skill2 "atlas-cac/character/skill"
   "atlas-cac/monster"
   "atlas-cac/skill"
   "atlas-cac/skill/information"
   "errors"
   "github.com/sirupsen/logrus"
)

func ProcessAttack(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attacked byte, damaged byte, attackedAndDamaged byte, stance byte, direction byte, rangedDirection byte, charge uint32, display byte, ranged bool, magic bool, speed byte, allDamage map[uint32][]uint32, x int16, y int16) error {
   return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attacked byte, damaged byte, attackedAndDamaged byte, stance byte, direction byte, rangedDirection byte, charge uint32, display byte, ranged bool, magic bool, speed byte, allDamage map[uint32][]uint32, x int16, y int16) error {
      attackCount := uint32(1)

      if !ranged && !magic {
         CloseRangeAttack(l)(worldId, channelId, mapId, characterId, skillId, skillLevel, attackedAndDamaged, display, direction, stance, speed, allDamage)
      } else if ranged {
         // ranged
      } else if magic {
         adjustedCharge := int32(-1)
         if skill.Is(skillId, skill.EvanFireBreath, skill.EvanIceBreath, skill.FirePoisonArchMagicianBigBang, skill.IceLighteningArchMagicianBigBang, skill.BishopBigBang) {
            adjustedCharge = int32(charge)
         }
         MagicAttack(l)(worldId, channelId, mapId, characterId, skillId, skillLevel, attackedAndDamaged, display, direction, stance, speed, adjustedCharge, allDamage)
         e, ok := GetSkillEffect(l)(characterId, skillId)
         if !ok {
            l.Errorf("Unable to retrieve skill %d effect for character %d.", skillId, characterId)
            return errors.New("retrieving skill effect")
         }
         attackCount = e.AttackCount()
         if e.Cooldown() > 0 {
            // apply cooldown
         }
      }

      l.Debugf("Attack count %d.", attackCount)

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

func GetSkillEffect(l logrus.FieldLogger) func(characterId uint32, skillId uint32) (*information.Effect, bool) {
   return func(characterId uint32, skillId uint32) (*information.Effect, bool) {
      s, err := skill2.GetSkillForCharacter(l)(characterId, skillId)
      if err != nil {
         l.WithError(err).Errorf("Unable to retrieve skill %d information for character %d.", skillId, characterId)
         return nil, false
      }

      if s.Level() == 0 {
         return nil, false
      }

      i, err := information.GetById(l)(skillId)
      if err != nil {
         l.WithError(err).Errorf("Unable to retrieve information for skill %d.", skillId)
         return nil, false
      } else {
         return &i.Effects()[s.Level() - 1], true
      }
   }
}