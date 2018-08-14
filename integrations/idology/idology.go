package idology

import (
    // "fmt"
    "io"
    "os"
    "net/http"
    // "io/ioutil"
    "bytes"
    "encoding/json"
	"kyc/common"
)

// Checks the customer with the KYC provider and returns a boolean indicating whether user is approved.
func CheckCustomer(customer *common.UserData) bool {

    server_url := "https://web.idologylive.com"

    body := new(bytes.Buffer)

    json.NewEncoder(body).Encode(customer)


    resp, _ := http.Post(server_url, "application/json;charset=utf-8", body)

    defer resp.Body.Close()

    io.Copy(os.Stdout, resp.Body)

	return true
}
