package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Item struct {
	name  string
	price float32
}

var itemList = make([]Item, 0)

func main() {
	homepage()

	for {
		var opt uint8
		fmt.Print("\nEnter option: ")
		_, err := fmt.Scan(&opt)
		if err != nil {
			errorMessage()
		}

		if opt == 0 {
			// list all items
			header()
			fmt.Println(getItems())
		} else if opt == 1 {
			// search an item an return its price
			fmt.Print("Search item name: ")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')

			header()
			fmt.Println(getItem(strings.Trim(name, "\n")))
		} else if opt == 2 {
			// create a new item
			var price float32

			fmt.Print("Enter item name: ")
			reader := bufio.NewReader(os.Stdin)
			name, _ := reader.ReadString('\n')
			fmt.Print("Enter item price: ")
			fmt.Scan(&price)

			newItem := createItem(strings.Trim(name, "\n"), price)
			fmt.Printf("Added %v\n", newItem.name)

			header()
			fmt.Println(getItems())

		} else if opt == 3 {
			break
		} else {
			fmt.Println("Invalid option! Enter a valid option")
			continue
		}

	}
}
