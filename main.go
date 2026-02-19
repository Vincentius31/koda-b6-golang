package main

import (
	"encoding/json"
	"fmt"
	"koda-b6-golang/services"
	"koda-b6-golang/utils"
	"net/http"
	"os"
	"strings"
)

var foodList []utils.Food
var service = services.FoodService{}

func loadData() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Error fatal: %v\n", r)
			os.Exit(1)
		}
	}()

	resp, err := http.Get("https://raw.githubusercontent.com/Vincentius31/koda-b6-golang/refs/heads/main/data/data.json")
	if err != nil {
		panic("Gagal Mengambil Data!")
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&foodList)
	if err != nil {
		panic("Gagal Passing JSON!")
	}
}

func searchFood() {
	fmt.Print("Cari Nama Makanan: ")
	var key string
	fmt.Scanln(&key)

	fmt.Println("\n--- Hasil Pencarian ---")
	for _, f := range foodList {
		if strings.Contains(strings.ToLower(f.NamaMakanan), strings.ToLower(key)) {
			fmt.Printf("ID: %d | %s\n", f.ID, f.NamaMakanan)
		}
	}
}

func orderFood() {
	fmt.Println("\n--- Pilih Menu ---")
	for _, f := range foodList {
		fmt.Printf("%d. %s\n", f.ID, f.NamaMakanan)
	}

	fmt.Print("Masukkan ID: ")
	var id int
	fmt.Scanln(&id)

	var selected *utils.Food
	for i := range foodList {
		if foodList[i].ID == id {
			selected = &foodList[i]
			break
		}
	}

	if selected == nil {
		fmt.Println("Menu tidak ditemukan.")
		return
	}

	choices := make([]string, 0, len(selected.Pilihan))
	for i := 1; i <= len(selected.Pilihan); i++ {
		key := fmt.Sprintf("opsi%d", i)
		choices = append(choices, key)
	}

	fmt.Println("\nPilihan Variasi:")
	for i, key := range choices {
		priceKey := "harga" + strings.Title(key) 
		
		label := selected.Pilihan[key]
		harga := selected.HargaMakanan[priceKey]
		fmt.Printf("%d. %s (Rp%d)\n", i+1, label, harga)
	}

	fmt.Print("Pilih nomor variasi: ")
	var optIdx int
	fmt.Scanln(&optIdx)

	if optIdx < 1 || optIdx > len(choices) {
		fmt.Println("Opsi tidak tersedia.")
		return
	}

	finalKey := choices[optIdx-1]
	finalPriceKey := "harga" + strings.Title(finalKey)
	
	finalLabel := selected.Pilihan[finalKey]
	finalHarga := selected.HargaMakanan[finalPriceKey]

	fmt.Printf("Harga satuan: Rp%d\n", finalHarga)
	fmt.Print("Jumlah (Qty): ")
	var qty int
	fmt.Scanln(&qty)

	if qty <= 0 {
		fmt.Println("Jumlah minimal 1.")
		return
	}

	service.AddToCart(utils.CartItem{
		Nama:  selected.NamaMakanan,
		Opsi:  finalLabel,
		Harga: finalHarga,
		Qty:   qty,
	})

	fmt.Printf("✅ %d %s (%s) masuk keranjang!\n", qty, selected.NamaMakanan, finalLabel)
}

func main() {
	loadData()

	for {
		fmt.Println("\n========== GOLANG FOOD APP ==========")
		fmt.Println("1. Cari Makanan")
		fmt.Println("2. Tambah ke Keranjang")
		fmt.Println("3. Checkout")
		fmt.Println("4. History Transaksi")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			searchFood()
		case 2:
			orderFood()
		case 3:
			done := make(chan bool)
			go service.Checkout(done)
			<-done
		case 4:
			showHistory()
		case 5:
			fmt.Println("Sampai jumpa!")
			return
		}
	}
}

func showHistory() {
	if len(service.History) == 0 {
		fmt.Println("Belum ada riwayat.")
		return
	}
	for _, h := range service.History {
		fmt.Printf("[%s] Total: Rp%d\n", h.Tanggal, h.Total)
	}
}
