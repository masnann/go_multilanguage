package models

type ProductModels struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Quantity    int     `json:"quantity"`
}

type RequestListProduct struct {
	Language string `json:"language"`
}

type ProductCreateRequest struct {
	Name        map[string]string `json:"name"`
	Description map[string]string `json:"description"`
	Price       float64           `json:"price"`
	Quantity    int               `json:"quantity"`
}
