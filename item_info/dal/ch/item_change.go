package ch

type ItemChange struct {
	ItemId     int64  `db:"item_id"`
	Time       int64  `db:"time"`
	ItemBefore string `db:"item_before"`
	ItemAfter  string `db:"item_after"`
}

func (i *ItemChange) TableName() string {
	return "item_change"
}
