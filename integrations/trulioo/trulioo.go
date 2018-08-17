package trulioo

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

    username := conf.Username
    password := conf.Password
    apiUrl := fmt.Sprintf(conf.Trulioo.ApiUrl, "0ac8ccee-ab7a-495e-8b88-a6da1bdcb6ae")

	request, _ := http.NewRequest("GET", apiUrl, nil)
    request.SetBasicAuth(username, password)

    client := &http.Client{}
    response, _ := client.Do(request)
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", response.StatusCode)
    }

    bodyBytes, _ := ioutil.ReadAll(response.Body)

    return &result.Status{"Low", string(bodyBytes)}, nil
}
