package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Org info is a tool that prints out a list of paid Apigee orgs
// In C you cannot return pointers in Function as the stack will be cleared.

type Organization struct {
	DisplayName  string
	Environments []string
	Type         string
	Created      int `json:"createdAt"`
}

func jsonGet(endpoint string) ([]byte, error) {
	mgmt := os.Getenv("APIGEE_MGMT_URL")
	user := os.Getenv("USERID")
	password := os.Getenv("USERPASSWORD")

	if mgmt == "" {
		mgmt = "https://api.enterprise.apigee.com/v1"
	}

	if user == "" || password == "" {
		fmt.Printf("Username or Password empty\n")
		os.Exit(1)
	}

	request, _ := http.NewRequest("GET", mgmt+endpoint, nil)
	request.SetBasicAuth(user, password)
	request.Header.Add("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		fmt.Printf("GET - %s returned %s", mgmt, err.Error())
		return []byte{}, err
	}

	var resString []byte
	resString, err = ioutil.ReadAll(resp.Body)
	return resString, err
}

func main() {

	orgs, _ := jsonGet("/organizations")

	var organizations []string
	//& assign values to organizations variable

	json.Unmarshal(orgs, &organizations)
	//fmt.Print(organizations[0])
	counter := 0
	for _, org := range organizations {
		var o Organization
		orgResp, _ := jsonGet("/organizations/" + org)
		json.Unmarshal(orgResp, &o)
		if o.Type == "paid" {
			fmt.Printf("%s is a paid org\n", o.DisplayName)
			counter++
		}
		//fmt.Printf(string(orgResp))
		fmt.Printf("Org Name is %s and is created at %d\n", o.DisplayName, o.Created)
		fmt.Printf("Total Number of paid orgs are:%d\n", counter)
	}

}
