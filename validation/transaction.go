package validation

func CheckTransaction(price int, paid int) (bool, int) {
	if price > paid {
		return false, price - paid
	} else {
		return true, paid - price
	}
}