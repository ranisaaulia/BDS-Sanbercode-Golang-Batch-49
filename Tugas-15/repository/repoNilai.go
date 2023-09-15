package repository

import (
	"Tugas-15/config"
	"Tugas-15/models"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

const (
	table3 = "nilai"
)

func GetAllNilai(ctx context.Context) ([]models.Nilai, error) {
	var nilaiNilai []models.Nilai
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By ID", table3)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var nilai models.Nilai
		if err = rowQuery.Scan(&nilai.ID,
			&nilai.Indeks,
			&nilai.Score,
			&nilai.Mahasiswa_ID,
			&nilai.Mata_Kuliah_ID); err != nil {
			return nil, err
		}

		nilaiNilai = append(nilaiNilai, nilai)
	}
	return nilaiNilai, nil
}

func InsertNilai(ctx context.Context, nilai models.Nilai) error {
	db, err := config.MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	// Hitung indeks_nilai berdasarkan nilai
	var indeks_nilai string
	switch {
	case nilai.Score >= 80:
		indeks_nilai = "A"
	case nilai.Score >= 70:
		indeks_nilai = "B"
	case nilai.Score >= 60:
		indeks_nilai = "C"
	case nilai.Score >= 50:
		indeks_nilai = "D"
	default:
		indeks_nilai = "E"
	}

	queryText := fmt.Sprintf("INSERT INTO %v (indeks, score, mahasiswa_id, mata_kuliah_id) VALUES (?, ?, ?, ?)", table3)
	_, err = db.ExecContext(ctx, queryText, indeks_nilai, nilai.Score, nilai.Mahasiswa_ID, nilai.Mata_Kuliah_ID)
	if err != nil {
		return err
	}
	return nil
}

func UpdateNilai(ctx context.Context, nilai models.Nilai, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	// Hitung indeks_nilai berdasarkan nilai yang akan diupdate
	var indeks_nilai string
	switch {
	case nilai.Score >= 80:
		indeks_nilai = "A"
	case nilai.Score >= 70:
		indeks_nilai = "B"
	case nilai.Score >= 60:
		indeks_nilai = "C"
	case nilai.Score >= 50:
		indeks_nilai = "D"
	default:
		indeks_nilai = "E"
	}

	queryText := fmt.Sprintf("UPDATE %v SET indeks=?, score=?, mahasiswa_id=?, mata_kuliah_id=? WHERE id=?",
		table3,
	)

	_, err = db.ExecContext(ctx, queryText, indeks_nilai, nilai.Score, nilai.Mahasiswa_ID, nilai.Mata_Kuliah_ID, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteNilai(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table3, id)

	s, err := db.ExecContext(ctx, queryText)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	check, err := s.RowsAffected()
	fmt.Println(check)
	if check == 0 {
		return errors.New("id tidak ada")
	}

	if err != nil {
		fmt.Println(err.Error())
	}

	return nil
}
