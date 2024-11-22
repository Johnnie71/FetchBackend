package models

type Item struct {
	ShortDescription string `json:"shortDescription"`
	Price string `json:"price"`
}

type Reciept struct {
	Retailer    string `json:"retailer" validate:"required"`
	PurchaseDate string `json:"purchaseDate" validate:"required"`
	PurchaseTime string `json:"purchaseTime" validate:"required"`
	Items        []Item `json:"items" validate:"required,min=1"`
	Total        string `json:"total" validate:"required"`
}