package models

type Item struct {
	ID         int
	Name       string
	Stock      int
	Categories []ItemCategory
}
