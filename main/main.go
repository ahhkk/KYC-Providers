package main

import (
	"fmt"
	kyc_provider "kyc/integrations/idology"//"kyc/integrations/example"
	"kyc/common"
)

func main() {

	// Example customer
	customer := &common.UserData{
		FirstName: "John",
		LastName: "Doe",
	}

	// Check a user against the example KYC provider
	check := kyc_provider.CheckCustomer(customer)
	if check {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
