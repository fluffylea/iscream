package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	cone                                string
	firstScoopNoAdditionalScoops        string
	firstScoopWithAdditionalScoops      string
	additionalScoopNoAdditionalScoops   string
	additionalScoopWithAdditionalScoops string
)

func init() {
	loadFile := func(name string) string {
		file, err := os.ReadFile(name)
		if err != nil {
			log.Fatalf("Unable to load %s", name)
		}
		return string(file)
	}
	cone = loadFile("cone.txt")
	firstScoopNoAdditionalScoops = loadFile("first_scoop_no_additional_scoops.txt")
	firstScoopWithAdditionalScoops = loadFile("first_scoop_with_additional_scoops.txt")
	additionalScoopNoAdditionalScoops = loadFile("additional_scoop_no_additional_scoops.txt")
	additionalScoopWithAdditionalScoops = loadFile("additional_scoop_with_additional_scoops.txt")
}

func buildCone(amount int) string {
	switch {
	case amount == 1:
		return firstScoopNoAdditionalScoops +
			cone
	case amount >= 2 && amount <= 1000:
		return additionalScoopNoAdditionalScoops +
			strings.Repeat(additionalScoopWithAdditionalScoops, amount-2) +
			firstScoopWithAdditionalScoops +
			cone
	default:
		return ":("
	}
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	values := r.URL.Query()
	amountStr := values.Get("amount")
	if amountStr == "" {
		http.Error(w, "Missing query parameter: amount", http.StatusBadRequest)
		return
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		http.Error(w, "Invalid query parameter: amount. It must be an integer", http.StatusBadRequest)
		return
	}
	cone := buildCone(amount)
	_, _ = w.Write([]byte(cone))
}

func main() {
	http.HandleFunc("GET /", getHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		panic(err)
	}
}
