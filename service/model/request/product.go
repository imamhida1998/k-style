package request

type CreateProduct struct {
	Nama     string `json:"nama"`
	Kategori string `json:"kategori"`
	Harga    int    `json:"harga"`
}

type UpdateProduct struct {
	ProductId string `json:"productId"`
	Nama      string `json:"nama"`
	Kategori  string `json:"kategori"`
	Harga     int    `json:"harga"`
}
