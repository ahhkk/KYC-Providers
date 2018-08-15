package idology

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
	"kyc/common"
)

// Checks the customer with the KYC provider and returns a boolean indicating whether user is approved.
func CheckCustomer(customer *common.UserData) bool {

    username := "modulus.dev"
    password := "L8kKhQe9TjiwrM?s"

    apiUrl := "https://web.idologylive.com"
    path := "/api/idiq.svc"

    body := url.Values{}

    body.Set("username", username)
    body.Add("password", password)
    body.Add("firstName", customer.FirstName)
    body.Add("lastName", customer.LastName)
    body.Add("address", customer.Address)
    body.Add("city", customer.City)
    body.Add("state", customer.State)
    body.Add("zip", customer.Zip)

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = path
    urlStr := u.String()

    fmt.Println(urlStr)

    client := &http.Client{}

    resp, _ := http.NewRequest("POST", urlStr, strings.NewReader(body.Encode()))
    resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    r, _ := client.Do(resp)

    defer r.Body.Close()

    if ( r.StatusCode == http.StatusOK) {
        bodyBytes, _ := ioutil.ReadAll(r.Body)
        fmt.Println(string(bodyBytes))
    }
	return true
}
