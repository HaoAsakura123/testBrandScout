package logic

import (
	"testing"
)

func setup() {
	Quotes = []Quote{
		{Author: "Confucius", Quote: "Bla Bla Bla"},
		{Author: "Nikitius", Quote: "Ble Ble Ble"},
		{Author: "Pupusius", Quote: "Blu Blu Blu"},
	}
}

func TestAddQuotes(t *testing.T) {
	setup()
	tests := []struct {
		name  string
		quote Quote
		err   bool
		idx   int
	}{
		{
			name:  "correct use AddQuotes",
			quote: Quote{Quote: "rapapam", Author: "pararam"},
			err:   false,
			idx:   3,
		},
		{
			name:  "quote already exists in memory",
			quote: Quote{Author: "Confucius", Quote: "Bla Bla Bla"},
			err:   true,
			idx:   -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			idx, err := AddQuotes(tt.quote)

			if (err != nil) != tt.err {
				t.Errorf("AddQuotes() error = %v, expected error : %v", err, tt.err)
				return
			}

			if idx != tt.idx {
				t.Errorf("AddQuotes() error = unexpected idx, current idx = %d, need idx = %d", idx, tt.idx)
			}

		})
	}

}

func TestDeleteQuotes(t *testing.T) {
	tests := []struct {
		name  string
		quote Quote
		err   bool
		idx   int
	}{
		{
			name:  "correct use DeleteQuotesFirst",
			quote: Quote{Author: "Confucius", Quote: "Bla Bla Bla"},
			err:   false,
			idx:   0,
		},
		{
			name:  "correct use DeleteQuotesSecond",
			quote: Quote{Author: "Nikitius", Quote: "Ble Ble Ble"},
			err:   false,
			idx:   1,
		},
		{
			name:  "quote not in  expected range",
			quote: Quote{},
			err:   true,
			idx:   52,
		},
	}

	for _, tt := range tests {
		setup()
		t.Run(tt.name, func(t *testing.T) {
			quote, err := DeleteQuote(tt.idx)

			if (err != nil) != tt.err {
				t.Errorf("DeleteQuotes() error = %v, expected error : %v", err, tt.err)
				return
			}

			if quote.Author != tt.quote.Author || quote.Quote != tt.quote.Quote {
				t.Errorf("Deleted but that's not it, deleted quote: %s, expected quote: %s, deleted author: %s, expected author: %s", quote.Author, tt.quote.Author, quote.Quote, tt.quote.Quote)
			}
		})
	}
}

func TestRandomQuote(t *testing.T) {
	setup()
	tests := []struct {
		name string
		err  bool
	}{
		{
			name: "random test first",
			err:  false,
		},
		{
			name: "random test second",
			err:  false,
		},
		{
			name: "random test third",
			err:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := RandomQuote()
			if (err != nil) != tt.err {
				t.Errorf("RandomQuote() error = %v, expected error : %v", err, tt.err)
				return
			}
		})
	}
}

func TestGetQuotesByOptions(t *testing.T) {
	tests := []struct {
		name       string
		author     string
		num        int
		massQuotes []Quote
	}{
		{
			name:       "Using with options",
			author:     "Confucius",
			num:        1,
			massQuotes: []Quote{{Author: "Confucius", Quote: "Bla Bla Bla"}},
		},
		{
			name:       "Using without options",
			author:     "",
			num:        0,
			massQuotes: Quotes,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
			quotes := GetQuotesByOptions(tt.author, tt.num)
			if len(quotes) != len(tt.massQuotes) {
				t.Errorf("testGetQuotesByOptions() error = uncorrrect num of quotes, received num : %d, expected num %d", len(quotes), len(tt.massQuotes))
				return
			}

			for iter, quote := range quotes {
				if quote.Author != tt.massQuotes[iter].Author || quote.Quote != tt.massQuotes[iter].Quote {
					t.Errorf("testGetQuotesByOptions() error = not that one quote, received quote %s, expected quote %s, received author %s, expected author %s", quote.Quote, Quotes[iter].Quote, quote.Author, Quotes[iter].Author)
					return
				}
			}
		})
	}
}

func TestAll(t *testing.T) {
	t.Run("DeleteQuotes", TestDeleteQuotes)
	t.Run("AddQuotes", TestAddQuotes)
	t.Run("RandomQuote", TestRandomQuote)
	t.Run("GetQuotesByOptions", TestGetQuotesByOptions)
}
