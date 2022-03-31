package skill

import (
	"atlas-cac/model"
	"atlas-cac/rest/requests"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	"strconv"
)

func ByCharacterAndSkillModelProvider(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32) model.Provider[Model] {
	return func(characterId uint32, skillId uint32) model.Provider[Model] {
		return requests.Provider[attributes, Model](l, span)(requestSkill(characterId, skillId), makeModel)
	}
}

func GetSkillForCharacter(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32) (Model, error) {
	return func(characterId uint32, skillId uint32) (Model, error) {
		return ByCharacterAndSkillModelProvider(l, span)(characterId, skillId)()
	}
}

func ByCharacterModelProvider(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) model.SliceProvider[Model] {
	return func(characterId uint32) model.SliceProvider[Model] {
		return requests.SliceProvider[attributes, Model](l, span)(requestSkills(characterId), makeModel)
	}
}

func GetSkillsForCharacter(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) ([]Model, error) {
	return func(characterId uint32) ([]Model, error) {
		return ByCharacterModelProvider(l, span)(characterId)()
	}
}

func makeModel(body requests.DataBody[attributes]) (Model, error) {
	id, err := strconv.ParseUint(body.Id, 10, 32)
	if err != nil {
		return Model{}, err
	}
	attr := body.Attributes
	return NewModel(uint32(id), attr.Level, attr.MasterLevel, attr.Expiration, false, false), nil
}
