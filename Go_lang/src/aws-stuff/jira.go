package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Issue struct {
	Key string `json:"key"`
}

type Response struct {
	Issues []Issue `json:"issues"`
}

func jsonGet() ([]byte, error) {
	user_id := os.Getenv("JIRA_USER")
	user_password := os.Getenv("JIRA_PASS")
	JIRA_URL := os.Getenv("JIRA_URL")
	USER := os.Getenv("USERVARIABLE1")
	request, _ := http.NewRequest("GET", JIRA_URL+"/rest/api/2/search?jql=assignee="+USER+"\\u0040apigee.com%20AND%20status%3DInvestigate", nil)
	request.SetBasicAuth(user_id, user_password)
	request.Header.Add("Accept", "application/json")

	client := new(http.Client)
	resp, err := client.Do(request)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}
	var resString []byte
	resString, err = ioutil.ReadAll(resp.Body)
	return resString, err
}
func main() {
	var res Response
	raw, _ := jsonGet()
	json.Unmarshal(raw, &res)
	for _, i := range res.Issues {
		fmt.Println(i.Key)
	}
}
