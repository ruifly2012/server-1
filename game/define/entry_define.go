package define

type HeroEntry struct {
	TypeID    int32   `json:"type_id"`
	AttID     int32   `json:"att_id"`
	Quality   int32   `json:"quality"`
	SpellList []int32 `json:"spell_list"`
}

type ItemEntry struct {
	TypeID   int32 `json:"type_id"`
	ItemType int32 `json:"item_type"`
	Quality  int32 `json:"quality"`
	Price    int32 `json:"price"`
}