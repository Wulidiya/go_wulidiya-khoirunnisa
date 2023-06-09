package main

import "fmt"

type Kendaraan struct {
	TotalRoda       int
	KecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {
	m.TambahKecepatan(10)
}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {
	m.KecepatanPerJam += kecepatanBaru
}

func main() {
	mobilCepat := Mobil{}
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()
	mobilCepat.Berjalan()

	mobilLamban := Mobil{}
	mobilLamban.Berjalan()

	fmt.Println("Kecepatan mobil cepat:", mobilCepat.KecepatanPerJam)
	fmt.Println("Kecepatan mobil lamban:", mobilLamban.KecepatanPerJam)
}

// Perubahan yang dilakukan:
// - Menggunakan huruf kapital pada nama struct dan method agar dapat diakses dari luar package.
// -Memberikan nama yang lebih deskriptif pada field dan method.
// -Menyederhanakan penulisan TambahKecepatan dengan operator +=.
// -Menambahkan import package fmt untuk menggunakan fungsi Println dalam fungsi main.
// -Menambahkan output untuk mencetak kecepatan mobil cepat dan mobil lamban.
// -Menambahkan spasi pada beberapa bagian untuk meningkatkan keterbacaan.
