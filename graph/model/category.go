package model

type Category struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Products []*Product `json:"products"`
}