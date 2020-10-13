package domain

import "errors"

var (
	items = map[string]*Item{
		"b101": {ItemID: "b101", Category: "Bags and Pouches", Price: 894.65, Rating: 4.2},
	}
)

// GetItem ...
func GetItem(itemID string) (*Item, error) {
	item := items[itemID]
	if item == nil {
		return nil, errors.New("No Item with ItemId: " + itemID + "found")
	}

	return item, nil
}

// PostItem ...
func PostItem(itemID string, category string, price float64, rating float32) bool {
	itemStruct := &Item{
		ItemID:   itemID,
		Category: category,
		Price:    price,
		Rating:   rating,
	}

	items[itemID] = itemStruct
	if _, ok := items[itemID]; ok {
		return ok
	}
	return false
}
