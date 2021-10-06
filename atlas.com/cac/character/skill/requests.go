package skill

import (
	"atlas-cac/rest/requests"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	skillsByCharacter              = charactersResource + "%d/skills"
	skillByCharacter               = skillsByCharacter + "/%d"
)

func requestSkills(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32) (*dataContainer, error) {
	return func(characterId uint32) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(skillsByCharacter, characterId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}

func requestSkill(l logrus.FieldLogger, span opentracing.Span) func(characterId uint32, skillId uint32) (*dataContainer, error) {
	return func(characterId uint32, skillId uint32) (*dataContainer, error) {
		ar := &dataContainer{}
		err := requests.Get(l, span)(fmt.Sprintf(skillByCharacter, characterId, skillId), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
