package main

import (
	"Tugas-2/Tugas-9/public"
	"fmt"
)

func main() {
	// Menampilkan Segitiga Sama Sisi
	segitiga := public.SegitigaSamaSisi{Alas: 3, Tinggi: 5}
	fmt.Println("Segitiga Sama Sisi:")
	fmt.Printf("Luas: %d\n", segitiga.Luas())
	fmt.Printf("Keliling: %d\n", segitiga.Keliling())

	// Menampilkan Persegi Panjang
	persegi := public.PersegiPanjang{Panjang: 3, Lebar: 7}
	fmt.Println("\nPersegi Panjang:")
	fmt.Printf("Luas: %d\n", persegi.Luas())
	fmt.Printf("Keliling: %d\n", persegi.Keliling())

	// Menampilkan Tabung
	silinder := public.Tabung{JariJari: 2.5, Tinggi: 9}
	fmt.Println("\nTabung:")
	fmt.Printf("Volume: %.2f\n", silinder.Volume())
	fmt.Printf("Luas Permukaan: %.2f\n", silinder.LuasPermukaan())

	// Menampilkan Balok
	kubus := public.Balok{Panjang: 3, Lebar: 3, Tinggi: 3}
	fmt.Println("\nBalok:")
	fmt.Printf("Volume: %.2f\n", kubus.Volume())
	fmt.Printf("Luas Permukaan: %.2f\n", kubus.LuasPermukaan())

	SamsungGalaxy := public.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	// Menampilkan data ponsel
	fmt.Println()
	fmt.Println(SamsungGalaxy.Display())

	// Menampilkan Luas Persegi
	fmt.Println(public.LuasPersegi(4, true))
	fmt.Println(public.LuasPersegi(8, false))
	fmt.Println(public.LuasPersegi(0, true))
	fmt.Println(public.LuasPersegi(0, false))

	// Menampilkan hitung penjumlahan
	fmt.Println()
	prefix := "hasil penjumlahan dari"
	kumpulanAngkaPertama := []int{6, 8}
	kumpulanAngkaKedua := []int{12, 14}

	hasil := public.HitungPenjumlahan(prefix, kumpulanAngkaPertama, kumpulanAngkaKedua)

	fmt.Printf("%s %v + %v = %d\n", prefix, kumpulanAngkaPertama, kumpulanAngkaKedua, hasil)
}
