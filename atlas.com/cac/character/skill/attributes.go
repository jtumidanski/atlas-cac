package skill

import "atlas-cac/rest/response"

type dataContainer struct {
	data     response.DataSegment
	included response.DataSegment
}

type dataBody struct {
	Id         string     `json:"id"`
	Type       string     `json:"type"`
	Attributes attributes `json:"attributes"`
}

type attributes struct {
	Level       uint8 `json:"level"`
	MasterLevel uint8 `json:"masterLevel"`
	Expiration  int64 `json:"expiration"`
}

func (a *dataContainer) UnmarshalJSON(data []byte) error {
	d, i, err := response.UnmarshalRoot(data, response.MapperFunc(emptySkillData))
	if err != nil {
		return err
	}

	a.data = d
	a.included = i
	return nil
}

func (a *dataContainer) Data() *dataBody {
	if len(a.data) >= 1 {
		return a.data[0].(*dataBody)
	}
	return nil
}

func (a *dataContainer) DataList() []dataBody {
	var r = make([]dataBody, 0)
	for _, x := range a.data {
		r = append(r, *x.(*dataBody))
	}
	return r
}

func emptySkillData() interface{} {
	return &dataBody{}
}