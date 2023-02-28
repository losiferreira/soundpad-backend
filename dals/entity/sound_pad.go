package entity

import "github.com/uptrace/bun"

type SoundPad struct {
	bun.BaseModel `bun:"table:sound_pads,alias:sp"`

	Id      int64  `bun:"id,pk,autoincrement"`
	Name    string `bun:"name,notnull"`
	OwnerId int64
	Owner   User `bun:"rel:belongs-to,join:owner_id=id"`

	Sounds []*Sound `bun:"m2m:sound_pad_to_sound,join:sound_pad=sound"`
}

type SoundPadToSound struct {
	bun.BaseModel `bun:"table:sound_pad_to_sound,alias:sp"`

	SoundPadID int64     `bun:"sound_pad_id,pk"`
	SoundPad   *SoundPad `bun:"rel:belongs-to,join:sound_pad_id=id"`
	SoundId    int64     `bun:"sound_id,pk"`
	Sound      *Sound    `bun:"rel:belongs-to,join:sound_id=id"`
}
