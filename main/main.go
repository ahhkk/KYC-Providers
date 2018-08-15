package main

import (
	"fmt"
	kyc_provider "kyc/integrations/idology"//"kyc/integrations/example"
	"kyc/common"
)

// SSN: 555-66-7777
// DOB: 08/01/1940

func main() {

	/*
		Example customer for IDology provider
	*/
	customer := &common.UserData{
		FirstName: "John",
		LastName: "Bredenkamp",
		Address: "147 Brentwood Drive",
		City: "Nashville",
		State: "TN",
		Zip: "37214",
	}

	/*
		Example customer for Sumandsubstance provider
	*/
	// customer := &common.UserData{
	// 	Id: "5b7298530a975a1df03bdd13",
	// }

	// Check a user against the example KYC provider
	check := kyc_provider.CheckCustomer(customer)
	fmt.Println(check)
	// if check {
	// 	fmt.Println("true")
	// } else {
	// 	fmt.Println("false")
	// }
}
