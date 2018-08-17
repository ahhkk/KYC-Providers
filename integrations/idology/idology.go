package idology

import (
    "fmt"
    "net/http"
    "net/url"
    "io/ioutil"
    "strings"
	"kyc/common"
    "kyc/result"
    "kyc/config"
    xj "github.com/basgys/goxml2json"
)

// Checks the customer with the KYC provider
// Returns type of status value and status value which represents the customer's status
// Returns error status when http request fails
func CheckCustomer(customer *common.UserData) (*result.Status, error) {

    conf := config.LoadConfiguration()

    username := conf.Username
    password := conf.Password
    apiUrl := conf.Idology.ApiUrl

    body := url.Values{}
    body.Set("username", username)
    body.Add("password", password)
    body.Add("firstName", customer.FirstName)
    body.Add("lastName", customer.LastName)
    body.Add("address", customer.Address)
    body.Add("city", customer.City)
    body.Add("state", customer.State)
    body.Add("zip", customer.Zip)

    request, _ := http.NewRequest("POST", apiUrl, strings.NewReader(body.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    response, _ := client.Do(request)
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", response.StatusCode)
    }

    // convert XML response to JSON
    bodyBytes, _ := ioutil.ReadAll(response.Body)
    json, _ := xj.Convert( strings.NewReader(string(bodyBytes)) )

    return &result.Status{"High", json.String()}, nil
}
