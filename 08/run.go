package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Rate struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	fmt.Println("Quest: Currency converter")

	if help := flag.Bool("help", false, "display help"); len(os.Args) == 1 || *help {
		fmt.Println("Usage: --from [currency] --to [currency]")
		os.Exit(1)
	}

	fromPtr := flag.String("from", "EUR", "from currency")
	toPtr := flag.String("to", "USD", "to currency")
	flag.Parse()

	from := strings.ToUpper(*fromPtr)
	to := strings.ToUpper(*toPtr)

	fmt.Printf("Scanning exchange rates from [%s] to [%s]\n", from, to)

	url := fmt.Sprintf("http://api.fixer.io/latest?symbols=%s,%s", from, to)

	// Handling request
	request, _ := http.Get(url)
	defer request.Body.Close()

	var rates Rate
	body, _ := ioutil.ReadAll(request.Body)
	if err := json.Unmarshal(body, &rates); err != nil {
		fmt.Printf("Error while decompositing JSON [%s]\n", err)
	} else {
		exchangeRate := rates.Rates[from]
		if exchangeRate > 0 {
			fmt.Printf("Exchange course from [%s] to [%s] is [%f]\n", from, to, exchangeRate)
		} else {
			fmt.Printf("Can't find exchange rate from [%s] to [%s]\n", from, to)
		}
	}

}
