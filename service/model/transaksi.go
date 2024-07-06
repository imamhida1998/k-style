package model

import "time"

type Transaksi struct {
	Id          string     `json:"id"`
	UserId      string     `json:"userId"`
	ProductId   string     `json:"productId"`
	NamaProduct string     `json:"namaProduct"`
	Status      string     `json:"status"`
	Total       int        `json:"total"`
	CreatedAt   *time.Time `json:"createdAt"`
}
