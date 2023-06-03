package responses

import "encoding/json"

type BookResponse struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Price       json.Number `json:"price"`
	Description string      `json:"description"`
	Rating      int         `json:"rating"`
	Discount    int         `json:"discount"`
}
