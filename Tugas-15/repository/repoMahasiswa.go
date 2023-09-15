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
	table1         = "mahasiswa"
	layoutDateTime = "2006-01-02 15:04:05"
)

func GetAllMahasiswa(ctx context.Context) ([]models.Mahasiswa, error) {
	var mhsiswa []models.Mahasiswa
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By ID", table1)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var mahasiswa models.Mahasiswa
		if err = rowQuery.Scan(&mahasiswa.ID,
			&mahasiswa.Nama); err != nil {
			return nil, err
		}

		mhsiswa = append(mhsiswa, mahasiswa)
	}
	return mhsiswa, nil
}

func InsertMahasiswa(ctx context.Context, mahasiswa models.Mahasiswa) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama) values('%v')", table1,
		mahasiswa.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMahasiswa(ctx context.Context, mahasiswa models.Mahasiswa, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET nama=? WHERE id=?",
		table1,
	)

	_, err = db.ExecContext(ctx, queryText, mahasiswa.Nama, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMahasiswa(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table1, id)

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
