package sumandsubstance

import (
    "fmt"
    "net/http"
    "io/ioutil"
	"kyc/common"
    "kyc/result"
    "kyc/config"
)

// Checks the customer with the KYC provider
// Returns type of status value and status value which represents the customer's status
// Returns error status when http request fails
func CheckCustomer(customer *common.UserData) (*result.Status, error) {

    conf := config.LoadConfiguration()

    apiUrl := fmt.Sprintf(conf.Sumandsubstance.ApiUrl, customer.Id)
    apiKey := conf.Sumandsubstance.ApiKey

    request, _ := http.NewRequest("GET", apiUrl, nil)
    request.Header.Add("Content-Type", "application/json")
    query := request.URL.Query()
    query.Add("key",apiKey)
    request.URL.RawQuery = query.Encode()

    client := &http.Client{}
    response, _ := client.Do(request)
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", response.StatusCode)
    }

    bodyBytes, _ := ioutil.ReadAll(response.Body)

    return &result.Status{"Low", string(bodyBytes)}, nil
}
