package model

import "time"

type Product struct {
	Id        string     `json:"id"`
	Nama      string     `json:"nama"`
	Kategori  string     `json:"kategori"`
	Harga     int        `json:"harga"`
	CreatedAt *time.Time `json:"createdAt"`
	CreatedBy string     `json:"createdBy"`
	UpdatedAt *time.Time `json:"updatedAt"`
	UpdatedBy string     `json:"updatedBy"`
}
