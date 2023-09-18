package models

type BangunDatar struct {
	ID       int
	Jenis    string
	Sisi     int
	Panjang  int
	Lebar    int
	JariJari int
	Hitung   string
}

type BangunRuang struct {
	ID       int
	Jenis    string
	Sisi     int
	Panjang  int
	Lebar    int
	Tinggi   int
	JariJari int
	Hitung   string
}

type Database interface {
	InsertBangunDatar(bd BangunDatar) error
	InsertBangunRuang(br BangunRuang) error
	GetBangunDatarByID(id int) (BangunDatar, error)
	GetBangunRuangByID(id int) (BangunRuang, error)
}

type Calculator interface {
	HitungLuas() float64
	HitungKeliling() float64
	HitungVolume() float64
	HitungLuasPermukaan() float64
}
