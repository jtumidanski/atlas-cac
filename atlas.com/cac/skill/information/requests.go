package information

import (
	"atlas-cac/rest/requests"
	"fmt"
)

const (
	skillServicePrefix string = "/ms/sis/"
	skillService              = requests.BaseRequest + skillServicePrefix
	skillsResource            = skillService + "skills"
	skillResource             = skillsResource + "/%d"
)

func requestSkill(skillId uint32) (*dataContainer, error) {
	ar := &dataContainer{}
	err := requests.Get(fmt.Sprintf(skillResource, skillId), ar)
	if err != nil {
		return nil, err
	}
	return ar, nil
}
