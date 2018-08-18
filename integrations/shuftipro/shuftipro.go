package shuftipro

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "net/url"
    "encoding/hex"
    // "encoding/json"
    "crypto/sha256"
    "strings"
    "kyc/common"
    "kyc/result"
    "kyc/config"
)

// Checks the customer with the KYC provider
// Returns type of status value and status value which represents the customer's status
// Returns error status when http request fails
func CheckCustomer(customer *common.UserData) (*result.Status, error) {

    conf := config.LoadConfiguration()

    apiUrl := conf.Shuftipro.ApiUrl
    apiKey := conf.Shuftipro.ApiKey
    clientId := ""
    reference := ""

    // The SHA256 hash of CONCATENATE(status_code, message, reference, secret_key).
    rawData := clientId + reference + apiKey
    hash := sha256.New()
    hash.Write([]byte(rawData))
    md := hash.Sum(nil)
    mdStr := hex.EncodeToString(md) // result string

    body := url.Values{}
    body.Set("client_id", clientId)
    body.Add("reference", reference)
    body.Add("signature", mdStr)

    request, _ := http.NewRequest("POST", apiUrl, strings.NewReader(body.Encode()))
    request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    response, _ := client.Do(request)
    defer response.Body.Close()

    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("Status error: %v", response.StatusCode)
    }

    bodyBytes, _ := ioutil.ReadAll(response.Body)

    return &result.Status{"Low", string(bodyBytes)}, nil
}
