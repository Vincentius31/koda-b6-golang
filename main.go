package main

import (
	"encoding/json"
	"fmt"
	"koda-b6-golang/services"
	"koda-b6-golang/utils"
	"net/http"
	"os"
)

var foodList []utils.Food
var service = services.FoodService{}

func loadData(){
	defer func(){
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

func main(){
	
}