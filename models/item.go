package models

type Item struct {
	Name       string
	ValueTotal int
	ValueUnity string
}
type ItemsResult struct {
	ItemsTotal int
	Items      []Item
}
