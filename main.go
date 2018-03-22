package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"go-payments-api/entity"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/payment", CreatePayment).Methods("POST")
	router.HandleFunc("/payment/{id}", UpdatePayment).Methods("PUT")
	router.HandleFunc("/payment/{id}", RemovePayment).Methods("DELETE")
	router.HandleFunc("/payments", GetPayments).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", router))
}
func CreatePayment(w http.ResponseWriter, r *http.Request) {
	payment := entity.Payment{
		// TODO: build aggregation payment, which includes the payment details with charges, debtor, fx, beneficiary
	}

	err := entity.Store(payment)
	if err != nil {
		respondError(w, err)
	}

	respondWithJSON(w, http.StatusCreated, map[string]int{"id": payment.ID})
}

func GetPayments(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func RemovePayment(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func UpdatePayment(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		panic(err.Error())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondError(w http.ResponseWriter, err error) {
	respondWithJSON(w, http.StatusBadRequest, map[string]string{
		"error": err.Error(),
	})
}
