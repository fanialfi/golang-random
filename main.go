package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano()) // digunakna untuk penentuan nilai seed

	// rand.Int() untuk generate angka random dalam bentuk int
	fmt.Println("random ke 1", rand.Int())
	fmt.Println("random ke 2", rand.Int())
	fmt.Println("random ke 3", rand.Int())

	// function rand.Intn(n) untuk mendapatkan angka random dengan batas 0 sampai < n
	// contoh rand.Intn(10) akan menghasilkan angka random dari 0 - 9
	fmt.Println(rand.Intn(25))

	// random tipe data string
	runes := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	randomString := func(length int) string {
		b := make([]rune, length)

		// jika hasil dari range hanya ditampung oleh 1 variabel, maka isinya hanyalah value dari slice-nya
		for i := range b {
			b[i] = runes[rand.Intn(len(runes))]
		}
		return string(b)

	}
	fmt.Println(randomString(25))
}
