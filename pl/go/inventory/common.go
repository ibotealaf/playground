package main

import (
	"errors"
	"fmt"
	"strings"
)

func homepage() {
	fmt.Println("0 -> All item list")
	fmt.Println("1 -> Search item")
	fmt.Println("2 -> Create new item")
	fmt.Println("3 -> Logout")
}

func errorMessage() {
	var errMssg error
	errMssg = errors.New("Invalid input")
	fmt.Printf("%v", errMssg)
}

func header() {
	fmt.Println("  ITEM  ||  PRICE($)")
	fmt.Println(":--------------------:")
}

func getItems() string {
	var result string
	for _, item := range itemList {
		result += fmt.Sprintf("  %v  ||  %v \n", item.name, item.price)
	}

	return result
}

func createItem(name string, price float32) Item {
	item := Item{name, price}
	itemList = append(itemList, item)
	return itemList[len(itemList)-1]
}

func getItem(name string) string {
	var result string
	for _, item := range itemList {
		present := strings.Contains(strings.ToLower(item.name), strings.ToLower(name))

		if present {
			result += fmt.Sprintf("  %v  ||  %v \n", item.name, item.price)
		}
	}
	return result
}
