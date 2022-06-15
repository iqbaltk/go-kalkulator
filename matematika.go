package matematika

func CekGanjilGenap(number ...int) []string {
	var data []string
	for _, i := range number {
		if i%2 == 0 {
			data = append(data, "Genap")
		} else {
			data = append(data, "Ganjil")
		}
	}
	return data
}
