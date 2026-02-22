package main

import (
	"fmt"
	"gophers_1/input"
	"gophers_1/validation"
)

func main() {
	price, paid := input.InputTransaction()
	valid, diff := validation.CheckTransaction(price, paid)
	
	if valid {
		fmt.Println("[SISTEM]: Transaksi Berhasil. Kembalian Anda:", diff)
	} else {
		fmt.Println("[SISTEM]: Transaksi Ditolak! Uang kurang", diff)
	}
}