package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

type Ukuran struct {
	Panjang  float64
	Lebar    float64
	Tinggi   float64
	Luas     chan float64
	Keliling chan float64
	Volume   chan float64
}

func phones(phones []string) {
	var wg sync.WaitGroup

	for i, phone := range phones {
		wg.Add(1)
		go func(i int, phone string) {
			defer wg.Done()
			time.Sleep(time.Second * time.Duration(i))
			fmt.Printf("%d. %s\n", i+1, phone)
		}(i, phone)
	}

	wg.Wait()
}

var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

func getMovies(moviesChannel chan string, movies ...string) {
	for i, movie := range movies {
		moviesChannel <- fmt.Sprintf("%d. %s", i+1, movie)
		time.Sleep(time.Second) // Delay selama satu detik
	}
	close(moviesChannel)
}

func luasLingkaran(radius float64, result chan<- float64) {
	luas := math.Pi * math.Pow(radius, 2)
	result <- luas
}

func kelilingLingkaran(radius float64, result chan<- float64) {
	keliling := 2 * math.Pi * radius
	result <- keliling
}

func volumeTabung(radius, height float64, result chan<- float64) {
	volume := math.Pi * math.Pow(radius, 2) * height
	result <- volume
}

func (u Ukuran) HitungLuas() {
	luas := u.Panjang * u.Lebar
	u.Luas <- luas
}

func (u Ukuran) HitungKeliling() {
	keliling := 2 * (u.Panjang + u.Lebar)
	u.Keliling <- keliling
}

func (u Ukuran) HitungVolume() {
	volume := u.Panjang * u.Lebar * u.Tinggi
	u.Volume <- volume
}

func main() {
	printPhones := []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	phones(printPhones)

	fmt.Println()
	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}

	// nomor 3 dan 4 di execute bergantian karena terdapat duplicate wg

	// fmt.Println()
	// radii := []float64{8, 14, 20}
	// height := 10

	// var wg sync.WaitGroup
	// luasChannel := make(chan float64)
	// kelilingChannel := make(chan float64)
	// volumeChannel := make(chan float64)

	// for _, radius := range radii {
	// 	wg.Add(1)
	// 	go func(radius float64) {
	// 		defer wg.Done()
	// 		luasLingkaran(radius, luasChannel)
	// 		kelilingLingkaran(radius, kelilingChannel)
	// 		volumeTabung(radius, float64(height), volumeChannel)
	// 	}(radius)
	// }

	// go func() {
	// 	wg.Wait()
	// 	close(luasChannel)
	// 	close(kelilingChannel)
	// 	close(volumeChannel)
	// }()

	// for radius := range luasChannel {
	// 	area := <-luasChannel
	// 	circumference := <-kelilingChannel
	// 	volume := <-volumeChannel

	// 	fmt.Printf("Jari-jari: %.2f\n", radius)
	// 	fmt.Printf("Luas Lingkaran: %.2f\n", area)
	// 	fmt.Printf("Keliling Lingkaran: %.2f\n", circumference)
	// 	fmt.Printf("Volume Tabung: %.2f\n\n", volume)
	// }

	fmt.Println()
	luasChan := make(chan float64)
	kelilingChan := make(chan float64)
	volumeChan := make(chan float64)

	hasil := Ukuran{Panjang: 5, Lebar: 3, Tinggi: 4, Luas: luasChan, Keliling: kelilingChan, Volume: volumeChan}

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		hasil.HitungLuas()
	}()

	go func() {
		defer wg.Done()
		hasil.HitungKeliling()
	}()

	go func() {
		defer wg.Done()
		hasil.HitungVolume()
	}()

	go func() {
		wg.Wait()
		close(luasChan)
		close(kelilingChan)
		close(volumeChan)
	}()

	select {
	case luas := <-luasChan:
		fmt.Printf("Luas Persegi Panjang: %.2f\n", luas)
	}
	select {
	case keliling := <-kelilingChan:
		fmt.Printf("Keliling Persegi Panjang: %.2f\n", keliling)
	}
	select {
	case volume := <-volumeChan:
		fmt.Printf("Volume Balok: %.2f\n", volume)
	}

}
