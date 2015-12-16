package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Org struct {
	Name string   `json:"displayName"`
	Env  []string `json:"environments"`
}

func jsonGet() ([]byte, error) {
	user_id := os.Getenv("USERID")
	user_password := os.Getenv("USERPASSWORD")
	request, _ := http.NewRequest("GET", "https://api.enterprise.apigee.com/v1/o/jai", nil)
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
	raw, _ := jsonGet()
	//	var mapp map[string]interface{}
	//	var mapp map[interface{}]interface{}
	var mapp Org
	json.Unmarshal(raw, &mapp)
	for _, i := range mapp.Env {
		fmt.Println(i)
	}
}
