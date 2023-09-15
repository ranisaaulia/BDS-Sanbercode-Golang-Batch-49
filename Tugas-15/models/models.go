package models

type Mahasiswa struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}

type MataKuliah struct {
	ID   uint   `json:"id"`
	Nama string `json:"nama"`
}

type Nilai struct {
	ID             uint   `json:"id"`
	Indeks         string `json:"indeks"`
	Score          int    `json:"score"`
	Mahasiswa_ID   int    `json:"mahasiswa_id"`
	Mata_Kuliah_ID int    `json:"mata_kuliah_id"`
}
