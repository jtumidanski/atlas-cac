package skill

import (
	"atlas-cac/rest/requests"
	"fmt"
)

const (
	charactersServicePrefix string = "/ms/cos/"
	charactersService              = requests.BaseRequest + charactersServicePrefix
	charactersResource             = charactersService + "characters/"
	skillsByCharacter              = charactersResource + "%d/skills"
	skillByCharacter               = skillsByCharacter + "/%d"
)

func requestSkills(characterId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(skillsByCharacter, characterId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}

func requestSkill(characterId uint32, skillId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(skillByCharacter, characterId, skillId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
