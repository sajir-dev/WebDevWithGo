package services

import (
	"../domain"
)

// GetItem ...
func GetItem(itemID string) *domain.Item {
	item, _ := domain.GetItem(itemID)
	return item
}

// PostItem ...
func PostItem(itemID string, category string, price float64, rating float32) bool {
	status := domain.PostItem(itemID, category, price, rating)
	return status
}
