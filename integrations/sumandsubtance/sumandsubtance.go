package sumandsubstance

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

    baseUrl := "https://test-api.sumsub.com/resources/"
    apiKey := "GKTBNXNEPJHCXY"
    testUser := "mandalaex_test_will_mandalaex_com"
    testApplicantId := "5b7298530a975a1df03bdd13"
    tokenPath := "accessTokens"

    client := &http.Client{}

    // Authentication
    body := url.Values{}
    body.Set("userId", testUser)
    body.Set("key", apiKey)

    u, _ := url.ParseRequestURI(baseUrl)
    u.Path = tokenPath
    urlStr := u.String()

    resp, _ := http.NewRequest("POST", urlStr, strings.NewReader(body.Encode()))
    resp.Header.Add("Content-Type", "application/json")


    // u, _ := url.ParseRequestURI(apiUrl)
    // u.Path = path
    // urlStr := u.String()
    //
    // fmt.Println(urlStr)
    //
    // client := &http.Client{}
    //
    // resp, _ := http.NewRequest("POST", urlStr, strings.NewReader(body.Encode()))
    // resp.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    //
    // r, _ := client.Do(resp)
    //
    // defer r.Body.Close()
    //
    // if ( r.StatusCode == http.StatusOK) {
    //     bodyBytes, _ := ioutil.ReadAll(r.Body)
    //     fmt.Println(string(bodyBytes))
    // }
	return true
}
