package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func packHandler(w http.ResponseWriter, r *http.Request) {
	availablePacks := []int{5000, 2000, 1000, 500, 250}

	qtyStr := r.URL.Query().Get("quantity")

	if r.URL.Query().Has("packs") {
		packsStr := r.URL.Query().Get("packs")
		packsAsStrings := strings.Split(packsStr, ",")
		availablePacks = []int{}

		for _, pack := range packsAsStrings {
			packAmt, err := strconv.Atoi(pack)
			if err != nil {
				panic(err)
			}
			availablePacks = append(availablePacks, packAmt)
		}
	}

	fmt.Fprintln(w, "Welcome to the stock packer: ")
	fmt.Fprint(w, "Available pack sizes: ")

	lastPack := len(availablePacks) - 1

	for i, pack := range availablePacks {
		fmt.Fprint(w, pack)
		if i < lastPack {
			fmt.Fprint(w, ", ")
		}
	}

	fmt.Fprintln(w)

	// string to int
	qty, err := strconv.Atoi(qtyStr)
	if err != nil {
		panic(err)
	}
	packing := packItems(availablePacks, qty)

	fmt.Fprintf(w, "Quantity: %d\n", qty)

	for _, entry := range packing {
		fmt.Fprintf(w, "%d x %d items\n", entry.quantity, entry.pack_volume)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the stock packer")
	fmt.Fprintln(w, "Go to /pack?quantity=1234 and change quantity to equal the amount of stock you want to manage")
	fmt.Fprintln(w, "To change pack sizes set the packs query param: /pack?quantity=399&packs=100,75,50,20,10")
	fmt.Fprintln(w, "Packs must be passed in size order")

}

func main() {
	fmt.Println("Listening on port 80")
	http.HandleFunc("/pack", packHandler)
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe(":80", nil))
}
