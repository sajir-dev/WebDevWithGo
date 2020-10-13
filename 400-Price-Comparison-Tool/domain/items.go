package domain

// Item ...
type Item struct {
	ItemID   string  `json: "item_id"`
	Category string  `json: "category"`
	Price    float64 `json: "price"`
	Rating   float32 `json: "rating"`
}
