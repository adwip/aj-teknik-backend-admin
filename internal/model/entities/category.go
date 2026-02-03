package entities

import "time"

type Category struct {
	ID          int       `db:"id"`
	SecureID    string    `db:"secure_id"`
	Name        string    `db:"name"`
	Parent      string    `db:"parent"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
	Description string    `db:"description"`
}
