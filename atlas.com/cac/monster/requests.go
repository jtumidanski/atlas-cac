package monster

import (
	"atlas-cac/rest/requests"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	monsterRegistryServicePrefix string = "/ms/morg/"
	monsterRegistryService              = requests.BaseRequest + monsterRegistryServicePrefix
	monstersResource                    = monsterRegistryService + "monsters"
	monsterResource                     = monstersResource + "/%d"
)

func getById(l logrus.FieldLogger) func(id uint32) (*DataContainer, error) {
	return func(id uint32) (*DataContainer, error) {
		ar := &DataContainer{}
		err := requests.Get(l)(fmt.Sprintf(monsterResource, id), ar)
		if err != nil {
			return nil, err
		}
		return ar, nil
	}
}
