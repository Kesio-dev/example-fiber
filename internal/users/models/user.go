package models

import "database/sql"

type User struct {
	ID        string         `db:"id"`
	Password  string         `db:"password"`
	Email     string         `db:"email"`
	Username  string         `db:"username"`
	Balance   float64        `db:"balance"`
	CreatedAt string         `db:"created_at"`
	RefID     sql.NullString `db:"referrer_id"`
	Avatar    string         `db:"avatar"`
	Role      string         `db:"user_role"`
	Status    string         `db:"status"`
}

type UserJSON struct {
	ID        string         `json:"id"`
	Email     string         `json:"email"`
	Username  string         `json:"username"`
	Balance   float64        `json:"balance"`
	CreatedAt string         `json:"created_at"`
	RefID     sql.NullString `json:"referrer_id"`
	RefCount  int            `json:"ref_count"`
	Avatar    string         `json:"avatar"`
	Role      string         `json:"user_role"`
	Status    string         `json:"status"`
}

type UserCreate struct {
	Password string  `db:"password"`
	Email    string  `db:"email"`
	Balance  float64 `db:"balance"`
}
