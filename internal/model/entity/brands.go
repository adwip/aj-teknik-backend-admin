package entity

type Brands struct {
	ID          string `db:"id"`
	SecureID    string `db:"secure_id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}
