package entity

import (
	"github.com/uptrace/bun"
)

type Sound struct {
	bun.BaseModel `bun:"table:sound,alias:s"`

	Id       int64  `bun:"id,pk,autoincrement"`
	Name     string `bun:"name"`
	FileName string `bun:"file_name"`
}
