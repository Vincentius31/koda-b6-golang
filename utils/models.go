package utils

type Food struct {
	ID           int               `json:"id"`
	NamaMakanan  string            `json:"namaMakanan"`
	Pilihan      map[string]string `json:"pilihan"`
	HargaMakanan map[string]int    `json:"hargaMakanan"`
}

type CartItem struct {
	Nama  string
	Opsi  string
	Harga int
	Qty   int
}

type Transaction struct {
	Tanggal string
	Items   []CartItem
	Total   int
}