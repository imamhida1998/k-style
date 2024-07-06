package response

type CreateTransaksi struct {
	TransaksiId string `json:"transaksiId"`
	Total       int    `json:"total"`
	// Amount      int    `json:"amount"`
}

type AcceptPayment struct {
	Message string
}
