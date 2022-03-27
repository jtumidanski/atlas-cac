package monster

import (
	"atlas-cac/rest/requests"
	"fmt"
)

const (
	monsterRegistryServicePrefix string = "/ms/morg/"
	monsterRegistryService              = requests.BaseRequest + monsterRegistryServicePrefix
	monstersResource                    = monsterRegistryService + "monsters"
	monsterResource                     = monstersResource + "/%d"
)

func getById(id uint32) requests.Request[attributes] {
	return requests.MakeGetRequest[attributes](fmt.Sprintf(monsterResource, id))
}
