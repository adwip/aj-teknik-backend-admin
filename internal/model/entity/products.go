package entity

import "time"

type Products struct {
	ID          string     `db:"id"`
	SecureID    string     `db:"secure_id"`
	Name        string     `db:"name"`
	Price       float64    `db:"price"`
	BrandId     string     `db:"brand_id"`
	Code        string     `db:"code"`
	Stock       int        `db:"stock"`
	CategoryId  int        `db:"category_id"`
	Description string     `db:"description"`
	Image       string     `db:"placeholder_image_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
