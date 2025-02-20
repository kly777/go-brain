package model

import "github.com/uptrace/bun"

// User is a struct that represents a user in the database
type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID	 int64  `bun:",pk,autoincrement"`
	Name string
	Password string
}

type Thing struct {
	bun.BaseModel `bun:"table:things,alias:t"`
	ID	 int64  `bun:",pk,autoincrement"`
	Name string
}

