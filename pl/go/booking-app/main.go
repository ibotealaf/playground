package main

import "fmt"

const concertName = "2023 EOY Party"
const totalTickets = 40
const ticketCost float32 = 99.99

var customers = make([]Customer, 0)
var availableTickets uint8 = 40

type Customer struct {
	firstName    string
	lastName     string
	email        string
	numOfTickets uint8
}

func main() {
	welcomeMessage()

	for availableTickets > 0 {
		customer := getUserDetails()
		isValidName, isValidEmail, isValidQty := validateUserDetails(customer)

		if isValidName && isValidEmail && isValidQty {
			if customer.numOfTickets > availableTickets {
				fmt.Printf("Sorry we have only %v tickets left", availableTickets)
				continue
			} else if customer.numOfTickets == availableTickets {
				fmt.Println("LUCKY YOU :) - You're getting our last set of tickets")
			}
			confirmPurchase(customer)
			go sendTicketInvoice(customer)

		} else {
			if !isValidName {
				fmt.Println("First name and last name must be at least 2 characters long")
			}
			if !isValidEmail {
				fmt.Println("Email is invalid")
			}
			if !isValidQty {
				fmt.Println("You must purchase at least 1 ticket")
			}
		}

		fmt.Printf("%v tickets already sold and we have %v tickets left. \n", (totalTickets - availableTickets), availableTickets)
	}

	stats(customers)
}
