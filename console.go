package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func consoleApp() {
	availablePacks := []int{5000, 2000, 1000, 500, 250}
	fmt.Println("Welcome to the stock packer: ")
	fmt.Print("Available pack sizes: ")
	for _, pack := range availablePacks {
		fmt.Print(pack)
		fmt.Print(", ")
	}
	fmt.Println()
	fmt.Println("To exit enter 'q'")
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter number of items to pack: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		if text == "q" {
			break
		}
		// string to int
		qty, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}
		packing := packItems(availablePacks, qty)

		for _, entry := range packing {
			fmt.Printf("%d x %d items\n", entry.quantity, entry.pack_volume)
		}
	}

}
