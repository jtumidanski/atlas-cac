package monster

import (
	"atlas-cac/rest/response"
	"encoding/json"
)

type DataContainer struct {
	data     response.DataSegment
	included response.DataSegment
}

type DataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes Attributes `json:"attributes"`
}

type DamageEntry struct {
	CharacterId uint32 `json:"characterId"`
	Damage      int64  `json:"damage"`
}

type Attributes struct {
	WorldId            byte          `json:"worldId"`
	ChannelId          byte          `json:"channelId"`
	MapId              uint32        `json:"mapId"`
	MonsterId          uint32        `json:"monsterId"`
	ControlCharacterId uint32        `json:"controlCharacterId"`
	X                  int16         `json:"x"`
	Y                  int16         `json:"y"`
	FH                 int16         `json:"fh"`
	Stance             byte          `json:"stance"`
	Team               int8          `json:"team"`
	MaxHp              int           `json:"maxHp"`
	Hp                 int           `json:"hp"`
	MaxMp              int           `json:"maxMp"`
	Mp                 int           `json:"mp"`
	DamageEntries      []DamageEntry `json:"damageEntries"`
}

func (c *DataContainer) MarshalJSON() ([]byte, error) {
	t := struct {
		Data     interface{} `json:"data"`
		Included interface{} `json:"included"`
	}{}
	if len(c.data) == 1 {
		t.Data = c.data[0]
	} else {
		t.Data = c.data
	}
	return json.Marshal(t)
}

func (c *DataContainer) UnmarshalJSON(data []byte) error {
	d, _, err := response.UnmarshalRoot(data, response.MapperFunc(EmptyMonsterData))
	if err != nil {
		return err
	}
	c.data = d
	return nil
}

func (c *DataContainer) Data() *DataBody {
	if len(c.data) >= 1 {
		return c.data[0].(*DataBody)
	}
	return nil
}

func (c *DataContainer) DataList() []DataBody {
	var r = make([]DataBody, 0)
	for _, x := range c.data {
		r = append(r, *x.(*DataBody))
	}
	return r
}

func EmptyMonsterData() interface{} {
	return &DataBody{}
}
