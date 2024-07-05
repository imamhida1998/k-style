package request

type CreateTransaksi struct {
	Id        string `json:"id"`
	UserId    string `json:"userId"`
	ProductId string `json:"productId"`
	Status    string `json:"status"`
	Quantity  int    `json:"quantity"`
}

type PaymentTransaksi struct {
	TransaksiId string `json:"transaksiId"`
	Amount      int    `json:"amount"`
}

type CancelTransaksi struct {
	TransaksiId string `json:"transaksiId"`
}
