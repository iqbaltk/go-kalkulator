package matematika

func CekGanjilGenap(number int) string {
	if number%2 == 0 {
		return "Genap"
	} else {
		return "Ganjil"
	}
}
