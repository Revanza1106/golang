package router

import (
	"go-postgres-crud/controller"
	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/produk",controller.AmbilSemuaProduk).Methods("GET","OPTIONS")
	router.HandleFunc("/api/produk/{id}",controller.AmbilProduk).Methods("GET","OPTIONS")
    router.HandleFunc("/api/produk",controller.TmbhProduk).Methods("POST","OPTIONS")
	router.HandleFunc("/api/produk/{id}",controller.UpdateProduk).Methods("PUT","OPTIONS")
	router.HandleFunc("/api/produk/{id}",controller.HapusProduk).Methods("PUT","OPTIONS")
    return router
}
