package config

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration config
type Configuration struct {
	LBAddr  string `json:"lb_addr"`
	OssAddr string `json:"oss_addr"`
}

var configuration *Configuration

func init() {
	file, _ := os.Open("./conf.json")
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration = &Configuration{}

	err := decoder.Decode(configuration)
	if err != nil {
		log.Printf("Error during package config decoder.Decode: %v", err)
		panic(err)
	}
}

// GetLBAddr get lb address
func GetLBAddr() string {
	return configuration.LBAddr
}

// GetOssAddr get oss address
func GetOssAddr() string {
	return configuration.OssAddr
}
