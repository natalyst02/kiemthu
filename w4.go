package kiemthu

func GetPrice(price int) string {
	s := ""
	if price < 0 || price > 120 {
		s = "Wrong age. Try again"
	} else if price >= 0 && price <= 9 {
		s = "Free"
	} else if price >= 10 && price <= 19 {
		s = "30000"
	} else if price >= 20 && price <= 59 {
		s = "50000"
	} else {
		s = "30000"
	}
	return s
}

func main() {
}
