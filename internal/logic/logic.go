package logic

import (
	"fmt"
	"math/rand"
	"net/http"
)

type Quote struct {
	Quote  string `json:"quote" binding:"required"`
	Author string `json:"author" binding:"required"`
}

// Так как по заданию не требуется особых условий, а именно разрешено хранить в памяти,
// то логику работы с БД я реализовал по сути здесь аналог /storage

var (
	Quotes = make([]Quote, 0)
)

func AddQuotes(quote Quote) (int, error) { //coveraged tests
	for _, elem := range Quotes {
		if elem.Author == quote.Author && elem.Quote == quote.Quote {
			return -1, fmt.Errorf("quote already exists")
		}
	}
	Quotes = append(Quotes, quote)
	return len(Quotes) - 1, nil
}

func DeleteQuote(id int) (Quote, error) { //coveraged tests
	var quote Quote
	if id < 0 || id > len(Quotes)-1 {
		return quote, fmt.Errorf("cannot delete element with id: %d", id)
	}
	quote = Quotes[id]
	Quotes = append(Quotes[:id], Quotes[id+1:]...)
	return quote, nil
}

func RandomQuote() (Quote, error) { //coveraged tests
	var quote Quote
	if len(Quotes) == 0 {
		return quote, fmt.Errorf("quote not exists")
	}
	randID := rand.Intn(len(Quotes))
	if randID >= len(Quotes) || randID < 0 {
		return quote, fmt.Errorf("number of quotes: %d", len(Quotes))
	}
	return Quotes[randID], nil
}

func GetQuotesByOptions(author string, num int) []Quote { //coveraged tests
	result := make([]Quote, 0)
	flag := false
	if num == 0 { // ну, если нет опций то любые записи подходят
		flag = true
	}
	for _, quote := range Quotes {
		if quote.Author == author || flag {
			result = append(result, quote)
		}
	}
	return result
}

func GetAuthors(r *http.Request) (string, int) {
	nums := 0
	params := r.FormValue("author")
	if len(params) > 0 {
		nums++
	}
	return params, nums
}
