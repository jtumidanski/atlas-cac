package character

import (
	skill2 "atlas-cac/character/skill"
	"atlas-cac/job"
	"atlas-cac/model"
	"atlas-cac/monster"
	"atlas-cac/rest/requests"
	"atlas-cac/skill"
	"atlas-cac/skill/information"
	"errors"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"math"
	"math/rand"
	"strconv"
)

func ProcessAttack(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attacked byte, damaged byte, attackedAndDamaged byte, stance byte, direction byte, rangedDirection byte, charge uint32, display byte, ranged bool, magic bool, speed byte, allDamage map[uint32][]uint32, x int16, y int16) error {
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, skillLevel byte, attacked byte, damaged byte, attackedAndDamaged byte, stance byte, direction byte, rangedDirection byte, charge uint32, display byte, ranged bool, magic bool, speed byte, allDamage map[uint32][]uint32, x int16, y int16) error {
		//TODO skillLevel is not a real value.
		attackCount := uint32(1)

		attackEffect := &information.Effect{}
		if skillId > 0 {
			ok := false
			attackEffect, ok = GetSkillEffect(l, span)(characterId, skillId)
			if !ok {
				return errors.New("cannot locate effect for skill being used")
			}
		}

		if !ranged && !magic {
			emitCloseRangeAttack(l, span)(worldId, channelId, mapId, characterId, skillId, skillLevel, attackedAndDamaged, display, direction, stance, speed, allDamage)
		} else if ranged {
			// ranged
		} else if magic {
			adjustedCharge := int32(-1)
			if skill.Is(skillId, skill.EvanFireBreath, skill.EvanIceBreath, skill.FirePoisonArchMagicianBigBang, skill.IceLighteningArchMagicianBigBang, skill.BishopBigBang) {
				adjustedCharge = int32(charge)
			}
			emitMagicAttack(l, span)(worldId, channelId, mapId, characterId, skillId, skillLevel, attackedAndDamaged, display, direction, stance, speed, adjustedCharge, allDamage)
			attackCount = attackEffect.AttackCount()
			if attackEffect.Cooldown() > 0 {
				// apply cooldown
			}
		}

		processMPChange(l, span)(worldId, channelId, mapId, characterId, skillId, attackEffect)

		l.Debugf("Attack count %d.", attackCount)

		for k, v := range allDamage {
			m, err := monster.GetById(l, span)(k)
			if err != nil {
				l.WithError(err).Errorf("Cannot locate monster %d which the attack from %d hit.", k, characterId)
				continue
			}
			totalDamage := uint32(0)
			for _, e := range v {
				totalDamage += e
			}
			monster.Damage(l, span)(worldId, channelId, mapId, m.UniqueId(), characterId, totalDamage)
		}

		if magic {
			IfHasSkill(l, span)(characterId, processMPEater(l, span)(worldId, channelId, mapId, characterId, allDamage), skill.FirePoisonWizardMPEater, skill.IceLightningWizardMPEater, skill.ClericMPEater)
		}
		return nil
	}
}

func processMPChange(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, effect *information.Effect) {
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, effect *information.Effect) {
		c, err := GetCharacterById(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate character %d who used skill %d.", characterId, skillId)
			return
		}

		mpChange := int16(0)
		if effect.MP() != 0 {
			mpChange += int16(effect.MP())
		}
		if effect.MPR() != 0 {
			mpChange += int16(effect.MPR() * float64(c.HP()))
		}
		if effect.MPCon() != 0 {
			mod := 1.0
			skillId := uint32(0)
			if job.IsInBranch(c.JobId(), job.FirePoisonMagician) {
				skillId = skill.FirePoisonMagicianElementAmplification
			} else if job.IsInBranch(c.JobId(), job.IceLightningMagician) {
				skillId = skill.IceLigthningMagicianElementAmplification
			} else if job.IsInBranch(c.JobId(), job.BlazeWizard2) {
				skillId = skill.BlazeWizardElementAmplification
			} else if job.IsInBranch(c.JobId(), job.Evan7) {
				skillId = skill.EvanMagicAmplification
			}
			if skillId != 0 {
				e, ok := GetSkillEffect(l, span)(characterId, skillId)
				if ok {
					mod = float64(e.X() / 100.0)
				}
			}
			mpChange -= int16(effect.MPCon()) * int16(mod)
			//TODO  account for infinity and concentrate
		}
		emitManaAdjustment(l, span)(characterId, mpChange)
	}
}

func processMPEater(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, allDamage map[uint32][]uint32) func(skillId uint32, effect *information.Effect) {
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, allDamage map[uint32][]uint32) func(skillId uint32, effect *information.Effect) {
		return func(skillId uint32, effect *information.Effect) {
			for mobId := range allDamage {
				applyMPEater(l, span)(worldId, channelId, mapId, characterId, skillId, mobId, effect)
			}
		}
	}
}

func applyMPEater(l logrus.FieldLogger, span opentracing.Span) func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, mobId uint32, effect *information.Effect) {
	return func(worldId byte, channelId byte, mapId uint32, characterId uint32, skillId uint32, mobId uint32, effect *information.Effect) {
		success := false
		if effect.Prop() == 1.0 {
			success = true
		} else if rand.Float64() < effect.Prop() {
			success = true
		}

		if !success {
			l.Debugf("Applied MP Eater for character %d with rate %d, but it failed.", characterId, effect.Prop())
			return
		}

		//TODO determine if mob is boss, skip if not
		m, err := monster.GetById(l, span)(mobId)
		if err != nil {
			l.WithError(err).Errorf("Unable to locate monster to apply MP Eater to.")
			return
		}
		mp := int32(math.Min(float64(m.MaxMP())*(float64(effect.X())/100.0), float64(m.MP())))
		if mp <= 0 {
			l.Debugf("No MP to be absorbed from monster %d.", mobId)
			return
		}

		l.Debugf("Applying MP Eater for character %d attack. They gained %d mana as a result.", characterId, mp)
		//TODO lower monster mana
		emitManaAdjustment(l, span)(characterId, int16(mp))
		showMPEater(l, span)(worldId, channelId, mapId, characterId, skillId)
	}
}

func IfHasSkill(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, exec func(skillId uint32, effect *information.Effect), skillIds ...uint32) {
	return func(characterId uint32, exec func(skillId uint32, effect *information.Effect), skillIds ...uint32) {
		if len(skillIds) == 0 {
			return
		}

		skills, err := skill2.GetSkillsForCharacter(l, span)(characterId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve skills for character %d.", characterId)
			return
		}

		var skillId uint32
		var effect *information.Effect
		for _, s := range skills {
			for _, sid := range skillIds {
				if s.Id() == sid {
					si, err := information.GetById(l, span)(sid)
					if err != nil {
						l.WithError(err).Errorf("Cannot retrieve effect for skill %d.", sid)
						return
					}
					skillId = sid
					effect = &si.Effects()[s.Level()-1]
				}
			}
		}
		if effect == nil {
			l.Debugf("Character %d does not have skills %s.", characterId, skillIds)
			return
		}

		exec(skillId, effect)
	}
}

func GetSkillEffect(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32) (*information.Effect, bool) {
	return func(characterId uint32, skillId uint32) (*information.Effect, bool) {
		s, err := skill2.GetSkillForCharacter(l, span)(characterId, skillId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve skill %d information for character %d.", skillId, characterId)
			return nil, false
		}

		if s.Level() == 0 {
			return nil, false
		}

		return GetSkillEffectWithLevel(l, span)(skillId, s.Level())
	}
}

func GetSkillEffectWithLevel(l logrus.FieldLogger, span opentracing.Span) func(skillId uint32, skillLevel uint8) (*information.Effect, bool) {
	return func(skillId uint32, skillLevel uint8) (*information.Effect, bool) {
		i, err := information.GetById(l, span)(skillId)
		if err != nil {
			l.WithError(err).Errorf("Unable to retrieve information for skill %d.", skillId)
			return nil, false
		} else {
			return &i.Effects()[skillLevel-1], true
		}
	}
}

func ByIdModelProvider(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) model.Provider[Model] {
	return func(characterId uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestCharacter(characterId), makeModel)
	}
}

func GetCharacterById(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (Model, error) {
	return func(characterId uint32) (Model, error) {
		return ByIdModelProvider(l, span)(characterId)()
	}
}

func makeModel(ca requests.DataBody[attributes]) (Model, error) {
	cid, err := strconv.ParseUint(ca.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	att := ca.Attributes
	r := Model{
		id:    uint32(cid),
		jobId: att.JobId,
		hp:    att.Hp,
	}
	return r, nil
}
