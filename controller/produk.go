package controller

import (
	"encoding/json" //package untuk merubah struct ke json dan sebaliknya
	"fmt"
	"strconv" // mengubah string menjadi tipe int

	"log"
	"net/http" //untuk mengakses request response di api

	"go-postgres-crud/models" //dimana produk di definisikan

	"github.com/gorilla/mux" // digunakan mendapatkan parameter dari router
	_ "github.com/lib/pq"    //postgres driver
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message.omitempty"`
}

type Response struct {
	Status  int             `json:"status"`
	Message string          `json:message`
	Data    []models.Produk `json:data`
}

// tambah produk
func TmbhProduk(w http.ResponseWriter, r *http.Request) {

	//variabel empaty dengan models.produk
	var produk models.Produk

	//decode data json ke produk
	err := json.NewDecoder(r.Body).Decode(&produk)

	if err != nil {
		log.Fatalf("Tidak bisa mencode dari request body. %v", err)
	}

	//memangil model lalu insert produk
	insertID := models.TambahProduk(produk)

	//format respon object
	res := response{
		ID:      insertID,
		Message: "produk ditambahkan",
	}
	json.NewEncoder(w).Encode(res)
}

// ambil satu produk dengan id
func AmbilProduk(w http.ResponseWriter, r *http.Request) {
	// set Header
	w.Header().Set("Context-type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Controler-Allow-Origin", "*")

	// mendapatkan id produk dari parametar request, keynya "id"
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("jika merubah sting ke int maka akan error. %v", err)
	}

	// memanggil models satu produk denga parameter id
	produk, err := models.AmbilSatuProduk(int64(id))

	if err != nil {
		log.Fatalf("tidak bisa mengambil data buku. %v", err)
	}
	json.NewEncoder(w).Encode(produk)
}

func AmbilSemuaProduk(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "Application/x-www-form-urlencoded")
	w.Header().Set("Access-Controler-Allow-Origin", "*")
	//memangil models semua produk
	produks, err := models.AmbilSemuaProduk()

	if err != nil {
		log.Fatalf("tidak bisa ambil data. %v", err)
	}
	var response Response
	response.Status = 1
	response.Message = "Sucess"
	response.Data = produks

	// kirim semua respon
	json.NewEncoder(w).Encode(response)
}

// update produk

func UpdateProduk(w http.ResponseWriter, r *http.Request) {

	// request parameter id
	params := mux.Vars(r)

	//konversi string menjadi int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("tidak bisa mengubah dari string ke int.  %v", err)
	}

	//variabel produk models.produk
	var produk models.Produk

	// decode json ke request variableproduk
	err = json.NewDecoder(r.Body).Decode(&produk)

	if err != nil {
		log.Fatalf("tidak bisa decode.  %v", err)
	}

	// pangil updateproduk untuk update data
	updateRows := models.UpdateProduk(int64(id), produk)

	// format message berupa string
	msg := fmt.Sprintf("Produk diperbarui. jumlah yang diperbarui %v rows/record", updateRows)

	// ini adalah format response message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	// kirim berupa respon
	json.NewEncoder(w).Encode(res)
}

// hapus produk
func HapusProduk(w http.ResponseWriter, r *http.Request) {

	//ambil request parameter id
	params := mux.Vars(r)

	//konversi string menjadi int
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Tidak bisa mengubah dari sting ke int. %v", err)
	}

	//pangil fungsi hapus produk,convert int ke int64
	deleteRows := models.HapusProduk(int64(id))

	//format message berupa string
	msg := fmt.Sprintf("Produk di hapus. Total data yang dihapus %v", deleteRows)

	// format respose message
	res := response{
		ID:      int64(id),
		Message: msg,
	}

	//sent
	json.NewEncoder(w).Encode(res)
}
