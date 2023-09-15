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
	table2 = "matakuliah"
)

func GetAllMatkul(ctx context.Context) ([]models.MataKuliah, error) {
	var matkul []models.MataKuliah
	db, err := config.MySQL()

	if err != nil {
		log.Fatal("Cant connect to MySQL", err)
	}

	queryText := fmt.Sprintf("SELECT * FROM %v Order By ID", table2)
	rowQuery, err := db.QueryContext(ctx, queryText)

	if err != nil {
		log.Fatal(err)
	}

	for rowQuery.Next() {
		var matakuliah models.MataKuliah
		if err = rowQuery.Scan(&matakuliah.ID,
			&matakuliah.Nama); err != nil {
			return nil, err
		}

		matkul = append(matkul, matakuliah)
	}
	return matkul, nil
}

func InsertMatkul(ctx context.Context, matakuliah models.MataKuliah) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("INSERT INTO %v (nama) values('%v')", table2,
		matakuliah.Nama)
	_, err = db.ExecContext(ctx, queryText)

	if err != nil {
		return err
	}
	return nil
}

func UpdateMatkul(ctx context.Context, matakuliah models.MataKuliah, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("UPDATE %v SET nama=? WHERE id=?",
		table2,
	)

	_, err = db.ExecContext(ctx, queryText, matakuliah.Nama, id)
	if err != nil {
		return err
	}

	return nil
}

func DeleteMatkul(ctx context.Context, id string) error {
	db, err := config.MySQL()
	if err != nil {
		log.Fatal("Can't connect to MySQL", err)
	}

	queryText := fmt.Sprintf("DELETE FROM %v where id = %s", table2, id)

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
