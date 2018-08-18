package config

import (
    "os"
    "fmt"
    "runtime"
    "encoding/json"
    "path/filepath"
)

type Api struct {
    ApiUrl string `json:"apiUrl"`
    ApiKey string `json:"apiKey"`
}

type Config struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Idology Api `json:"idology"`
    Sumandsubstance Api `json:"sumandsubstance"`
    Trulioo Api `json:"trulioo"`
    Shuftipro Api `json:"shuftipro"`
}

func LoadConfiguration() Config {

    var config Config
    file := "development.json"

    _, fn, _, _ := runtime.Caller(0)

    configFile, err := os.Open( filepath.Join(filepath.Dir(fn), file) )
    defer configFile.Close()

    if err != nil {
        fmt.Println(err.Error())
    }

    jsonParser := json.NewDecoder(configFile)
    jsonParser.Decode(&config)
    return config
}
