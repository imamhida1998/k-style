package request

type CreateTransaksi struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type PaymentTransaksi struct {
	TransaksiId string `json:"transaksiId"`
	Amount      int    `json:"amount"`
}

type CancelTransaksi struct {
	TransaksiId string `json:"transaksiId"`
}

type AcceptTransaksi struct {
	TransaksiId string `json:"transaksiId"`
}
