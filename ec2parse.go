package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Records struct {
	Records []Record `json:"Records"`
}

type Record struct {
	Identity  Identity `json:"userIdentity"`
	EventTime string   `json:"eventTime"`
	EventName string   `json:"eventName"`
	Region    string   `json:"awsRegion"`
	IPAddress string   `json:"sourceIPAddress"`
	Agent     string   `json:"userAgent"`
}

type Identity struct {
	UserType  string `json:"type"`
	Resource  string `json:"arn"`
	AccessKey string `json:"accessKeyId"`
	AccountID string `json:"accountId"`
	Username  string `json:"userName"`
}

func main() {
	Logs, err := os.Open("traillog.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("File has been opened.")

	defer Logs.Close() // waits for all other functions to run before executing

	byteValue, _ := ioutil.ReadAll(Logs)

	var records Records

	json.Unmarshal(byteValue, &records)

	// iterating through instance

	for i := 0; i < len(records.Records); i++ {
		fmt.Println("Identity: " + records.Records[i].Identity.UserType)
		fmt.Println("Resource: " + records.Records[i].Identity.Resource)
		fmt.Println("Access Key: " + records.Records[i].Identity.AccessKey)
		fmt.Println("Account ID: " + records.Records[i].Identity.AccountID)
		fmt.Println("EventTime: " + records.Records[i].EventTime)
		fmt.Println("EventName: " + records.Records[i].EventName)
		fmt.Println("Region: " + records.Records[i].Region)
		fmt.Println("IP Address: " + records.Records[i].IPAddress)
		fmt.Println("Agent: " + records.Records[i].Agent)

	}

}
