package main

import (
	"fmt"
	"task1/input"
	"task1/validation"
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