package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"testbrandscout/internal/logic"
)

func HandleQuotes(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodPost:
		QuotesPost(w, r)

	case http.MethodGet:
		QuotesGET(w, r)

	default:
		log.Println("INFO: Method not allowed")
		http.Error(w, "cannnot use this method", http.StatusMethodNotAllowed)
		return

	}
}

func HandleDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		log.Println("INFO: Method not allowed")
		http.Error(w, "cannnot use this method", http.StatusMethodNotAllowed)
		return
	}

	path := r.URL.Path
	id, err := strconv.Atoi(path[len("/quotes/"):])
	if err != nil {
		log.Println("INFO: cannot convert request")
		http.Error(w, "Plese check your request", http.StatusBadRequest)
		return
	}

	quote, err := logic.DeleteQuote(id)
	if err != nil {
		log.Printf("INFO: uncorrect idx: %v", err)
		http.Error(w, "INFO: uncorrect idx", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status":       "success",
		"deletedID":    id,
		"deletedQuote": quote.Quote,
		"author":       quote.Author,
	}); err != nil {
		log.Printf("INFO: JSON encode error: %v", err)
	}

}

func HandleRandom(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		log.Println("INFO: method not allowed")
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	quote, err := logic.RandomQuote()

	if err != nil {
		log.Println("INFO: Not exists any quote")
		http.Error(w, "not exists any quote", http.StatusNoContent)
		return
	}

	w.Header().Set("content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "access",
		"quote":  quote.Quote,
		"author": quote.Author,
	}); err != nil {
		log.Printf("INFO: JSON encode error: %v", err)
	}
}

func QuotesPost(w http.ResponseWriter, r *http.Request) {

	var quote logic.Quote
	if err := json.NewDecoder(r.Body).Decode(&quote); err != nil {
		log.Printf("INFO: cant deserialisation data: %v", err)
		http.Error(w, "Plese check your request", http.StatusBadRequest)
		return
	}

	idx, err := logic.AddQuotes(quote)
	if err != nil {
		log.Printf("INFO: cant add quote to quotes: %v", err)
		http.Error(w, "quote already exists", http.StatusConflict)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"quote":  quote.Quote,
		"author": quote.Author,
		"idx":    idx,
	}); err != nil {
		log.Printf("DEBUG: JSON encode error: %v", err)
	}

}

func QuotesGET(w http.ResponseWriter, r *http.Request) {
	author, num := logic.GetAuthors(r)
	quotes := logic.GetQuotesByOptions(author, num)

	if len(quotes) == 0 {
		log.Println("INFO: quotes not exists")
		http.Error(w, "Content not exists", http.StatusNoContent)
		return
	}

	if err := json.NewEncoder(w).Encode(map[string]interface{}{
		"status": "success",
		"quotes": quotes,
	}); err != nil {
		log.Printf("DEBUG: JSON encode error: %v", err)
	}

}
