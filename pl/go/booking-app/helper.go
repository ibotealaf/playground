package main

import (
	"fmt"
	"math/rand"
	"slices"
	"strconv"
	"strings"
	"time"
)

func welcomeMessage() {
	fmt.Println("#############################################")
	fmt.Printf("Welcome to the %v booking page.\n", concertName)
	fmt.Printf("We have total of %v up for grabs and %v remaining tickets\n", totalTickets, availableTickets)
	fmt.Println("#############################################")
	fmt.Println("Proceed to buy ticket")
}

func sendTicketInvoice(user Customer) {
	time.Sleep(10 * time.Second)
	for i := 0; i < int(user.numOfTickets); i++ {
		fmt.Println("\n######################")
		fmt.Printf("NAME: %v %v\n", user.firstName, user.lastName)
		fmt.Printf("TICKETS ID: %v\n", rand.Uint32())
		fmt.Printf("AMOUNT: #%v\n", ticketCost)
		fmt.Println("######################")
	}
}

func validateUserDetails(customer Customer) (bool, bool, bool) {
	isValidName := len(customer.firstName) >= 2 && len(customer.lastName) >= 2
	isValidEmail := strings.Contains(customer.email, "@")
	isValidQty := customer.numOfTickets > 0

	return isValidName, isValidEmail, isValidQty
}

func getUserDetails() Customer {
	var firstName string
	var lastName string
	var email string
	var numOfTickets uint8

	fmt.Print("\nEnter your first name: ")
	fmt.Scan(&firstName)
	fmt.Print("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Print("Enter email address: ")
	fmt.Scan(&email)
	fmt.Print("How many tickets do you need: ")
	fmt.Scan(&numOfTickets)

	customer := Customer{firstName, lastName, email, numOfTickets}

	return customer
}

func confirmPurchase(customer Customer) {
	costOfQty := ticketCost * float32(customer.numOfTickets)
	var confirm string

	fmt.Printf("Hey %v %v the cost for %v ticket is $%v\n", customer.firstName, customer.lastName, customer.numOfTickets, costOfQty)
	fmt.Println("Proceed to make payment...")
	fmt.Println("Press yes/y to confirm and no/n to cancel")
	fmt.Scan(&confirm)

	confirm = strings.ToLower(confirm)
	if confirm == "y" || confirm == "yes" {
		fmt.Printf("Thanks for your purchase %v. Sent invoice to %v\n", customer.firstName, customer.email)
		availableTickets -= customer.numOfTickets

		customers = append(customers, customer)
	} else {
		fmt.Println("Operation terminated!")
	}
}

func stats(customers []Customer) string {
	namesOfCustomers := ""
	customersFirstNames := []string{}

	for _, customer := range customers {
		present := slices.Contains(customersFirstNames, customer.firstName)

		if !present {
			customersFirstNames = append(customersFirstNames, customer.firstName)
		}
	}

	filteredFirstNames := len(customersFirstNames)
	if filteredFirstNames >= 1 {
		namesOfCustomers += customersFirstNames[0]

		for i := 1; i < filteredFirstNames; i++ {
			if i == filteredFirstNames-1 {
				namesOfCustomers += " & "
			} else {
				namesOfCustomers += ", "
			}
			namesOfCustomers += customersFirstNames[i]
		}
	} else {
		return "No customers yet"
	}

	statement := "We had a total of " + strconv.Itoa(filteredFirstNames) + " customer(s) with the following names: " + namesOfCustomers
	return statement
}
