package main

import (
	"api/config"
	"api/modules/accounts"
	"api/modules/transactions"
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
		db.AutoMigrate(&transactions.Transaction{})
	}

	//setup transaction module
	transactionRepo := transactions.NewRepo(db)
	transactionService := transactions.NewService(&transactionRepo)
	transactionHandler := transactions.NewHandler(&transactionService)

	//setup account module
	accountRepo := accounts.NewRepo()
	accountService := accounts.NewService(&accountRepo)
	accountHandler := accounts.NewHandler(&accountService)

	router := mux.NewRouter()

	var port = ":" + os.Getenv("API_APP_PORT")
	router.HandleFunc("/validate", accountHandler.Validate).Methods("POST")
	router.HandleFunc("/disbursement", transactionHandler.Disbursement).Methods("POST")
	router.HandleFunc("/disbursement-callback", transactionHandler.DisbursementCallback).Methods("PUT")
	log.Println("listening at port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
