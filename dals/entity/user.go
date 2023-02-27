package entity

import "github.com/uptrace/bun"

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	Id    int64  `bun:"id,pk,autoincrement"`
	Name  string `bun:"name,notnull"`
	Email string `bun:"email,notnull"`
}
