package product

type Product struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	ImageUrl    string `json:"image_url"`
	IsPurchased bool   `json:"is_purchased"`
}

type DeleteRequest struct {
	Id     string `json:"id" binding:"required"`
	UserId string `json:"user_id" binding:"required"`
}

type CreateProductRequest struct {
	Name        string `json:"name" binding:"required"`
	UserId      string `json:"user_id" binding:"required"`
	Description string `json:"description"`
	Link        string `json:"link"`
	EventId     string `json:"event_id" binding:"required"`
}

type ProductDetailsResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	UserId      string `json:"user_id"`
	ImageUrl    string `json:"image_url"`
	Description string `json:"description"`
	Link        string `json:"link"`
	EventId     string `json:"event_id"`
	IsPurchased bool   `json:"is_purchased"`
}

type EditRequest struct {
	Id          string `json:"id" binding:"required"`
	EventId     string `json:"event_id" binding:"required"`
	Name        string `json:"name,omitempty"`
	UserId      string `json:"user_id" binding:"required"`
	ImageUrl    string `json:"image_url,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
}

type UpdateProductPurchaseStatus struct {
	ProductId   string `json:"product_id" binding:"required"`
	EventId     string `json:"event_id" binding:"required"`
	IsPurchased bool   `json:"is_purchased" binding:"required"`
}
