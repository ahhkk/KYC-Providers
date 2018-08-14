package main

import (
	"fmt"
	kyc_provider "kyc/integrations/idology"//"kyc/integrations/example"
	"kyc/common"
)

// SSN: 555-66-7777
// DOB: 08/01/1940

func main() {

	// Example customer
	customer := &common.UserData{
		Username: "modulus.dev",
		Password: "L8kKhQe9TjiwrM?s",
		FirstName: "John",
		LastName: "Bredenkamp",
		Address: "147 Brentwood Drive",
		City: "Nashville",
		State: "TN",
		Zip: "37214",
	}

	// Check a user against the example KYC provider
	check := kyc_provider.CheckCustomer(customer)
	if check {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
