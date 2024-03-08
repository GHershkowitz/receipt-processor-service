package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"net/http"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []*Item `json:"items"`
	Total        string  `json:"total"`
}

type Item struct {
	Description string `json:"shortDescription"`
	Price       string `json:"price"`
}

var dbReceipts = make(map[string]*Receipt)
var dbReceiptScore = make(map[string]int)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/receipts/process", ProcessReceiptsHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", GetPointsHandler).Methods("GET")

	http.Handle("/", r)
	fmt.Println("server has started...")
	http.ListenAndServe(":8080", nil)
}

func ProcessReceiptsHandler(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		fmt.Println("error decoding", err)
		return
	}

	receiptID := uuid.New().String()
	CreateReceipt(receiptID, receipt)
	response := map[string]string{"id": receiptID}
	jsonResponse(w, response)
}

func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	receiptID := vars["id"]

	if receiptID == "" {
		jsonResponse(w, "invalid receipt_id")
		return
	}

	// Check if points were already calculated.
	if points, ok := dbReceiptScore[receiptID]; ok {
		response := map[string]int{"points": points}
		jsonResponse(w, response)
		return
	}

	points := CalculateReceiptScore(dbReceipts[receiptID])
	dbReceiptScore[receiptID] = points

	response := map[string]int{"points": points}
	jsonResponse(w, response)
}

func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
