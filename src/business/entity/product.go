package entity

type Product struct {
	ID       int64  `db:"id" json:"id"`
	Name     string `db:"name" json:"name"`
	Price    int64  `db:"price" json:"price"`
	Size     int64  `db:"size" json:"size"`
	Quantity int64  `db:"quantity" json:"quantity"`
}

type ProductParams struct {
	ID   int64  `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}
