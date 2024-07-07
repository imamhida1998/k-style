package response

type CreateTransaksi struct {
	TransaksiId string `json:"transaksiId"`
	Total       int    `json:"total"`
	// Amount      int    `json:"amount"`
}

type AcceptPayment struct {
	Message string
}

type DetailTransaksi struct {
	NamaProduct     string `json:"namaProduct"`
	JumlahPembelian int    `json:"jumlahPembelian"`
	StatusTransaksi string `json:"statusTransaksi"`
	TotalHarga      int    `json:"totalHarga"`
}
