package main

import (
	"fmt"
	"testing"
)

func TestStats(t *testing.T) {
	customers := make([]Customer, 0)
	customer1 := Customer{"tobi", "brave", "tobi@brave.com", 1}
	customer2 := Customer{"Saint", "Ives", "st@ives.com", 1}
	customer3 := Customer{"Lauren", "Ives", "thelaw@hey.com", 1}

	customerFirstNameDup := Customer{"tobi", "Ives", "ives@brave.com", 2}

	gotWhenNoCustomer := stats(customers)
	expectedNoCustomer := "No customers yet"

	if expectedNoCustomer != gotWhenNoCustomer {
		fmt.Println("WHEN THERE ARE NO CUSTOMERS")
		t.Errorf("expected %q \ngot %q", expectedNoCustomer, gotWhenNoCustomer)
	}

	customers = append(customers, customer1)
	gotWhenOnlyOneCustomer := stats(customers)
	expectedWhenOnlyOneCustomer := "We had a total of 1 customer(s) with the following names: tobi"

	if expectedWhenOnlyOneCustomer != gotWhenOnlyOneCustomer {
		fmt.Println("WHEN THERE'S JUST ONE CUSTOMER")
		t.Errorf("expected %q \ngot %q", expectedWhenOnlyOneCustomer, gotWhenOnlyOneCustomer)
	}

	customers = append(customers, customer2)
	gotWhenTwoCustomers := stats(customers)

	expectedWhenTwoCustomers := "We had a total of 2 customer(s) with the following names: tobi & Saint"

	if expectedWhenTwoCustomers != gotWhenTwoCustomers {
		fmt.Println("WHEN THERE ARE NO DUPLICATIONS IN CUSTOMERS NAMES")
		t.Errorf("expected %q \ngot %q", expectedWhenTwoCustomers, gotWhenTwoCustomers)
	}

	customers = append(customers, customerFirstNameDup)
	gotWhenDuplicateInCustomerFirstNames := stats(customers)
	expectedWhenDuplicateInCustomerFirstName := "We had a total of 2 customer(s) with the following names: tobi & Saint"

	if expectedWhenDuplicateInCustomerFirstName != gotWhenDuplicateInCustomerFirstNames {
		fmt.Println("WHEN THERE ARe DUPLICATIONS IN CUSTOMERS NAMES")
		t.Errorf("expected %q \ngot %q", expectedWhenDuplicateInCustomerFirstName, gotWhenDuplicateInCustomerFirstNames)
	}

	customers = append(customers, customer3)
	gotWhenMultipleCustomers := stats(customers)
	expectedWhenMultipleCustomers := "We had a total of 3 customer(s) with the following names: tobi, Saint & Lauren"

	if expectedWhenMultipleCustomers != gotWhenMultipleCustomers {
		fmt.Println("WHEN THERE ARE MORE THAN TWO CUSTOMERS")
		t.Errorf("expected %q \ngot %q", expectedWhenMultipleCustomers, gotWhenMultipleCustomers)
	}
}

func TestValidateUserDetails(t *testing.T) {
	isValidNameResult, isValidEmailResult, isValidQtyResult := validateUserDetails(Customer{"tobi", "brave", "tobi@brave.com", 1})
	isValidNameExpectedResult := true
	isValidEmailExpectedResult := true
	isValidQtyExpectedResult := true

	if isValidNameExpectedResult != isValidNameResult {
		t.Errorf("Expected name validation to be %v\nName validation is %v", isValidNameExpectedResult, isValidNameResult)
	}
	if isValidEmailExpectedResult != isValidEmailResult {
		t.Errorf("Expected name validation to be %t\nName validation is %t", isValidEmailExpectedResult, isValidEmailResult)
	}
	if isValidQtyExpectedResult != isValidQtyResult {
		t.Errorf("Expected quantity validation to be %t\nQuantity validation is %t", isValidQtyExpectedResult, isValidQtyResult)
	}
}
