package main

import (
	"fmt"
	"math/rand"
)

type Pemain struct {
	Dadu int
}

func acakDadu(pemain Pemain) (angka []int, angka1 int, angka6 int) {
	for i := 0; i < pemain.Dadu; i++ {
		num := rand.Intn(6) + 1
		angka = append(angka, num)
		if num == 1 {
			angka1 += 1
		} else if num == 6 {
			angka6 += 6
		}
	}
	return
}

func main() {
	var N, M int
	fmt.Print("Masukkan Jumlah Pemain : ")
	fmt.Scanln(&N)
	fmt.Print("Masukkan Jumlah Dadu : ")
	fmt.Println(&M)

}
