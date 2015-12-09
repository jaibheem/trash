package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	//"reflect"
	//	"unsafe"
)

type Pods struct {
	Name   string `json:"name"`
	Region string `json:"region"`
}

type Expand struct {
	ExternalIp       string `json:"externalIP"`
	ExternalHostName string `json:"externalHostName"`
	InternalHostName string `json:"internalHostName"`
	InternalIp       string `json:"internalIP"`
}

func getRequests(endpoint string) ([]byte, error) {
	userid := os.Getenv("USERID")
	userpasswd := os.Getenv("USERPASSWORD")
	mgmt := os.Getenv("APIGEE_MGMT_URL")
	if mgmt == "" {
		mgmt = "https://api.enterprise.apigee.com/v1"
	}

	if userid == "" || userpasswd == "" {
		fmt.Printf("Username or Password empty\n")
		os.Exit(1)
	}
	req, _ := http.NewRequest("GET", mgmt+endpoint, nil)
	req.SetBasicAuth(userid, userpasswd)
	req.Header.Add("Accept", "application/json")
	client := new(http.Client)
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println(err, mgmt)
	}
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}

func listEnv() (environments []string) {
	envs, _ := getRequests("/o/jai/e")
	json.Unmarshal(envs, &environments)
	return environments
	//	for _, env := range environments {
	//		var mp []string
	//		mps, _ := getRequests("/o/jai/e/" + env + "/servers")
	//		json.Unmarshal(mps, &mp)
	//		for _, i := range mp {
	//			fmt.Printf(env)
	//			fmt.Printf("\t")
	//			fmt.Printf(i)
	//			fmt.Printf("\t")
	//		}
	//	}
	//	fmt.Printf("\n")
}

func main() {
	envs := listEnv()
	var wg sync.WaitGroup
	for _, env := range envs {
		wg.Add(1)
		go func(env string) {
			defer wg.Done()
			var mps []string
			servers, _ := getRequests("/o/jai/e/" + env + "/servers")
			json.Unmarshal(servers, &mps)
			var wg1 sync.WaitGroup
			for _, mp := range mps {

			}
		}(env)
	}
	wg.Wait()
}
