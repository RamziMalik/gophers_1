package input

import "fmt"

func InputTransaction() (int, int){
	var price, paid int
	fmt.Print("Harga Barang: ")
	fmt.Scan(&price)
	fmt.Print("Uang Pembeli: ")
	fmt.Scan(&paid)
	return price, paid
}