# KYC Package

## Integrations

The main package will initiate a KYC check with a chosen KYC provider.

The KYC provider is changed by changing the package being imported to the _kyc_provider_ named import.

To create a new integration, simply add a new subpackage under integrations.

An example integration package is provided. As you can see, all an integration package must do is provide a single exported method:

```go
func CheckCustomer(customer *common.UserData) bool
```

The main package will call this method in a goroutine to perform a check. The method should block until a response is has been received from the KYC provider.

## Update

A single exported method changed to:

```go
func CheckCustomer(customer *common.UserData) (*result.Status, error) - returns error when http request fails
```

## Packages
go get github.com/basgys/goxml2json
go get golang.org/x/text/encoding


## Config Variables
Test user, apiUrl and apiKey for each provider are saved in "/config/development.json"
