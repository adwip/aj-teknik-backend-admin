package entities

import "time"

type Products struct {
	ID          int64      `db:"id"`
	SecureID    string     `db:"secure_id"`
	Name        string     `db:"name"`
	PriceId     string     `db:"price_id"`
	BrandId     string     `db:"brand_id"`
	Code        string     `db:"code"`
	Stock       float64    `db:"stock"`
	CategoryId  string     `db:"category_id"`
	Description string     `db:"description"`
	Image       string     `db:"placeholder_image_id"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
}
