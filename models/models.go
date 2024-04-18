package models

import (
	"database/sql"
	"fmt"
	"go-postgres-crud/config"
	"log"

	_ "github.com/lib/pq" //golang driver
)

type Produk struct {
	ID          int64  `json:"id"`
	Nama_produk string `json:"Nama_produk"`
	Jenis       string `json:"Jenis"`
	exp         string `json:"exp"`
}

func TambahProduk(produk Produk) int64 {

	// mengkoneksikan ke db postgres
	db := config.CreateConnection()

	//tutup koneksi
	defer db.Close()

	// insert query
	sqlStatement := `INSERT INTO Produk(Nama_produk,Jenis,exp) VALUES ($1,$2,$3)`

	//id yang dimasukin akan disimpan di id ini
	var id int64

	//Scan function akan menyimpan insert id didalam id
	err := db.QueryRow(sqlStatement, produk.Nama_produk, produk.Jenis, produk.exp)

	if err != nil {
		log.Fatalf("tidak bisa eksekusi query.  %v", err)
	}

	// return
	return id
}

// ambil semua produk
func AmbilSemuaProduk() ([]Produk, error) {
	//   koneksi ke db
	db := config.CreateConnection()
	// tutup koneksi
	defer db.Close()

	var produks []Produk

	// SELECT query
	sqlSatement := `SELECT * FROM produk`

	// eksekusi query
	rows, err := db.Query(sqlSatement)

	if err != nil {
		log.Fatalf("tidak bisa eksekusi query.  %v", err)
	}

	//tutupeksekusi sql query
	defer rows.Close()

	//ambil datanya
	for rows.Next() {
		var produk Produk

		//ambil datanya unmashal ke struct
		err = rows.Scan(&produk.ID, &produk.Nama_produk, &produk.Jenis, &produk.exp)

		if err != nil {
			log.Fatalf("tidak bisa mengambil data.  %v", err)
		}
		//masukan slice produks
		produks = append(produks, produk)
	}
	return produks, err
}

// ambil satu prduk
func AmbilSatuProduk(id int64) (Produk, error) {
	db := config.CreateConnection()

	defer db.Close()

	var produk Produk

	//sql query
	sqlStatement := `SELECT * FROM produk WHERE	id=$1`

	//eksekusi sql statement
	row := db.QueryRow(sqlStatement, id)

	err := row.Scan(&produk.ID, &produk.Nama_produk, &produk.Jenis, &produk.exp)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("tidak ada yng dicari")
		return produk, nil
	case nil:
		return produk, nil
	default:
		log.Fatalf("tidak bisa ambil data.  %v", err)
	}

	return produk, err
}

// update produk di db
func UpdateProduk(id int64, produk Produk) int64 {

	//koneksi ke db
	db := config.CreateConnection()

	//tutup koneksi
	defer db.Close()

	//buat sql query
	sqlStatement := `UPDATE produk SET nama_produk=$2,jenis=$3,exp=$4 WHERE id=$1`

	//eksekusi sql satement
	res, err := db.Exec(sqlStatement, produk.ID, produk.Nama_produk, produk.Jenis, produk.exp)

	if err != nil {
		log.Fatalf("tidak dapat akses. %v", err)
	}

	//cek jumlah row/data di update
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak dapat mengecek dsta. %v", err)
	}
	fmt.Printf("total record update %v\n", rowsAffected)

	return rowsAffected
}

// hapus produk
func HapusProduk(id int64) int64 {
	db := config.CreateConnection()

	defer db.Close()
	sqlStatement := `DELETE FROM PRODUK WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatalf("tidak bisa akses query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("tidak bisa cari data. %v", err)
	}
	fmt.Printf("total data terhapus %v", rowsAffected)
	return rowsAffected
}
