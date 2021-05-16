package monster

func GetMonster(id uint32) (*Monster, error) {
   resp, err := getById(id)
   if err != nil {
      return nil, err
   }

   d := resp.Data()
   n := makeMonster(id, d.Attributes)
   return &n, nil
}

func makeMonster(id uint32, att MonsterAttributes) Monster {
   return NewMonster(id, att.ControlCharacterId, att.MonsterId, att.X, att.Y, att.Stance, att.FH, att.Team)
}
