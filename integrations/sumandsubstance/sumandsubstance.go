package sumandsubstance

import (
    // "fmt"
    "net/http"
    "io/ioutil"
	"kyc/common"
)

// Checks the customer with the KYC provider and returns a boolean indicating whether user is approved.
func CheckCustomer(customer *common.UserData) string {

    apiUrl := "https://test-api.sumsub.com/resources/applicants/" + customer.Id + "/status"
    apiKey := "GKTBNXNEPJHCXY"

    client := &http.Client{}
    req, _ := http.NewRequest("GET", apiUrl, nil)
    q := req.URL.Query()
    q.Add("key",apiKey)
    req.URL.RawQuery = q.Encode()

    req.Header.Add("Content-Type", "application/json")

    r, _ := client.Do(req)

    defer r.Body.Close()

    bodyBytes, _ := ioutil.ReadAll(r.Body)

	return string(bodyBytes)
}
