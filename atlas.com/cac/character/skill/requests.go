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

func requestSkills(characterId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(skillsByCharacter, characterId))
}

func requestSkill(characterId uint32, skillId uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(skillByCharacter, characterId, skillId))
}
