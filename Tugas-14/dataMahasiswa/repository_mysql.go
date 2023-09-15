package datamahasiswa

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"
	"tugas-14/config"
	"tugas-14/models"
)

const (
	table          = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

// GetAll
func GetAll(ctx context.Context) ([]models.Mahasiswa, error) {
	var mhsiswa []models.Mahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By created_at DESC", table)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mahasiswa models.Mahasiswa
		var createdAt, updatedAt string
		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.Nama,
			&mahasiswa.MataKuliah,
			&mahasiswa.IndeksNilai,
			&mahasiswa.Nilai,
			&createdAt,
			&updatedAt); err != nil {
			return nil, err
		}

		//  Change format string to datetime for created_at and updated_at
		mahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		mahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)
		if err != nil {
			log.Fatal(err)
		}

		mhsiswa = append(mhsiswa, mahasiswa)
	}
	return mhsiswa, nil
}

// Insert Movie
func InsertMahasiswa(ctx context.Context, mahasiswa models.Mahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		return err
	}
	defer db.Close()

	// Hitung indeks_nilai berdasarkan nilai
	var indeksNilai string
	switch {
	case mahasiswa.Nilai >= 80:
		indeksNilai = "A"
	case mahasiswa.Nilai >= 70:
		indeksNilai = "B"
	case mahasiswa.Nilai >= 60:
		indeksNilai = "C"
	case mahasiswa.Nilai >= 50:
		indeksNilai = "D"
	default:
		indeksNilai = "E"
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama, mata_kuliah, indeks_nilai, nilai, created_at, updated_at) VALUES (?, ?, ?, ?, NOW(), NOW())", table)
	_, err = db.ExecContext(ctx, queryText, mahasiswa.Nama, mahasiswa.MataKuliah, indeksNilai, mahasiswa.Nilai)
	if err != nil {
		return err
	}
	return nil
}

func Update(ctx context.Context, mahasiswa models.Mahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	// Hitung indeks_nilai berdasarkan nilai yang akan diupdate
	var indeksNilai string
	switch {
	case mahasiswa.Nilai >= 80:
		indeksNilai = "A"
	case mahasiswa.Nilai >= 70:
		indeksNilai = "B"
	case mahasiswa.Nilai >= 60:
		indeksNilai = "C"
	case mahasiswa.Nilai >= 50:
		indeksNilai = "D"
	default:
		indeksNilai = "E"
	}

	queryText := fmt.Sprintf("UPDATE %v SET nama=?, mata_kuliah=?, indeks_nilai=?, nilai=?, updated_at=NOW() WHERE id=?",
		table,
	)

	_, err = db.ExecContext(ctx, queryText, mahasiswa.Nama, mahasiswa.MataKuliah, indeksNilai, mahasiswa.Nilai, id)
	if err != nil {
		return err
	}

	return nil
}

func Delete(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table, id)

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
