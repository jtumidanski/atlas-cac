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

func getById(id uint32) (*DataContainer, error) {
   ar := &DataContainer{}
   err := requests.Get(fmt.Sprintf(monsterResource, id), ar)
   if err != nil {
      return nil, err
   }
   return ar, nil
}
