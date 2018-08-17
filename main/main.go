package main

import (
	"fmt"
	kyc_provider "kyc/integrations/idology" // idology, sumandsubstance, trulioo, shuftipro,
	"kyc/common"
)

// SSN: 555-66-7777
// DOB: 08/01/1940

func main() {

	// Example customer for IDology provider
	customer := &common.UserData{
		Id: "5b7298530a975a1df03bdd13",
		FirstName: "John",
		LastName: "Smith",
		Address: "222333 PEACHTREE PLACE",
		City: "Atlanta",
		State: "GA",
		Zip: "30318",
	}

	// Example customer for Sumandsubstance provider
	// customer := &common.UserData{
	//	Id: "5b7298530a975a1df03bdd13",
	// }

	// Check a user against the KYC provider
	check, err := kyc_provider.CheckCustomer(customer)

	if err == nil {
		fmt.Printf("\nResponse Type: %s\n", check.Type)
		fmt.Printf("\nResponse Value: \n")
		fmt.Printf("\n%v\n", check.Value)
	} else {
		fmt.Println(err.Error())
	}
}
