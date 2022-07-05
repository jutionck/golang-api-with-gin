package model

type Product struct {
	ProductId   string `json:"productId" binding:"required"`
	ProductName string `json:"productName" binding:"required"`
	IsStatus    bool   `json:"isStatus"`
}
