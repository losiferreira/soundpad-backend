package entity

import "github.com/uptrace/bun"

type SoundPadToSound struct {
	bun.BaseModel `bun:"table:sound_pad_to_sound,alias:spts"`

	SoundPadId int64     `bun:"sound_pad_id,pk"`
	SoundPad   *SoundPad `bun:"rel:belongs-to,join:sound_pad_id=id"`
	SoundId    int64     `bun:"sound_id,pk"`
	Sound      *Sound    `bun:"rel:belongs-to,join:sound_id=id"`
}
