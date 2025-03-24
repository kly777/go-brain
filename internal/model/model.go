package model

import "database/sql"

// User 代表数据库中的用户
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password sql.NullString `json:"-"`
}

// Thing 代表数据库中的物品
type Thing struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

