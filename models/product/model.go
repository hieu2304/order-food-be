package product_model

type Product struct {
	ID       uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Category string  `json:"category"`
}

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
}
