package services

import (
	"koda-b6-golang/utils"
	"fmt"
	"time"
)

// Interface untuk standarisasi
type CartManager interface {
	AddToCart(item utils.CartItem)
	Checkout(done chan bool)
}

type FoodService struct {
	CartItems []utils.CartItem
	History   []utils.Transaction
}

func (s *FoodService) AddToCart(item utils.CartItem) {
	s.CartItems = append(s.CartItems, item)
}

func (s *FoodService) Checkout(done chan bool) {
	if len(s.CartItems) == 0 {
		fmt.Println("\n[!] Keranjang kosong!")
		done <- false
		return
	}

	total := 0
	for _, item := range s.CartItems {
		total += item.Harga * item.Qty
	}

	newTx := utils.Transaction{
		Tanggal: time.Now().Format("2006-01-02 15:04:05"),
		Items:   append([]utils.CartItem{}, s.CartItems...),
		Total:   total,
	}

	s.History = append(s.History, newTx)
	s.CartItems = []utils.CartItem{}
	
	fmt.Printf("\n[SUCCESS] Checkout berhasil! Total: Rp%d\n", total)
	done <- true
}