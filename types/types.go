package types

type Login struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type Response struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
type Registration struct {
	Name        string `json:"user_sname"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	PinCode     string `json:"pin_code"`
	Latitude    string `json:"latitude"`
	Longitude   string `json:"longitude"`
	Password    string `json:"password"`
	MapUrl      string `json:"map_url"`
}
type ProductListing struct {
	ProductName               string `json:"product_name"`
	ProductCategory           string `json:"product_category"`
	ProductPrice              string `json:"product_price"`
	ProductDescription        string `json:"product_description"`
	ProductQuantity           string `json:"product_quantity"`
	ProductStock              string `json:"product_stock"`
	ProductDiscountPercentage string `json:"product_discount_percentage"`
	ProductCollectionID       string `json:"product_collection_id"`
}
