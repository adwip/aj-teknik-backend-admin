package dtos

type ProductListResponseDto struct {
	SecureID    string  `db:"secure_id"`
	Name        string  `db:"name"`
	Code        string  `db:"code"`
	Price       float64 `db:"price"`
	Unit        string  `db:"unit"`
	Stock       float64 `db:"stock"`
	Brand       string  `db:"brand"`
	Category    string  `db:"category"`
	Image       string  `db:"image"`
	Description string  `db:"description"`
}
