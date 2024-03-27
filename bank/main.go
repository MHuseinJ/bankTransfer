package main

import (
	"bank/config"
	"bank/modules/accounts"
	"bank/modules/transactions"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load ENV
	godotenv.Load(".env")
	dbConfig := config.LoadDB()
	db := dbConfig.DB
	if os.Getenv("SKIP_MIGRATION") == "false" {
		db.AutoMigrate(&accounts.Account{})
	}

	//setup transaction module
	transactionService := transactions.NewService()
	transactionHandler := transactions.NewHandler(&transactionService)

	//setup account module
	accountRepo := accounts.NewRepo(db)
	accountService := accounts.NewService(&accountRepo)
	accountHandler := accounts.NewHandler(&accountService)

	router := mux.NewRouter()

	var port = ":" + os.Getenv("APP_PORT")
	router.HandleFunc("/bank/validate", accountHandler.Validate).Methods("POST")
	router.HandleFunc("/bank/setlement", transactionHandler.Settlement).Methods("POST")
	router.HandleFunc("/bank/transfer", transactionHandler.InitTransfer).Methods("POST")
	log.Println("listening at port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
