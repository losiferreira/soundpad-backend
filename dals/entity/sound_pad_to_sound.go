package entity

import "github.com/uptrace/bun"

type SoundPad struct {
	bun.BaseModel `bun:"table:sound_pads,alias:sp"`

	Id      int64  `bun:"id,pk,autoincrement"`
	Name    string `bun:"name,notnull"`
	OwnerId int64
	Owner   User `bun:"rel:belongs-to,join:owner_id=id"`

	Sounds []*Sound `bun:"m2m:sound_pad_to_sound,join:SoundPad=Sound"`
}
