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

func listEnv(wg *sync.WaitGroup) {
	var environments []string
	envs, _ := getRequests("/o/jai/e")
	json.Unmarshal(envs, &environments)
	for _, env := range environments {
		var mp []string
		mps, _ := getRequests("/o/jai/e/" + env + "/servers")
		json.Unmarshal(mps, &mp)
		for _, i := range mp {
			fmt.Printf(env)
			fmt.Printf("\t")
			fmt.Printf(i)
			fmt.Printf("\t")
		}
	}
	fmt.Printf("\n")
	wg.Done()
}

func listApiProducts(wg *sync.WaitGroup) {
	var apiproducts []string
	apis, _ := getRequests("/o/jai/apis")
	json.Unmarshal(apis, &apiproducts)
	for _, api := range apiproducts {
		fmt.Printf(api)
		fmt.Printf("\t")
	}
	fmt.Printf("\n")
	wg.Done()
}

func getMps(wg *sync.WaitGroup) {
	var p []Pods
	pod, _ := getRequests("/o/jai/pods")
	json.Unmarshal(pod, &p)
	for _, i := range p {
		fmt.Printf(i.Name)
	}
	wg.Done()
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

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go listEnv(&wg)
	wg.Add(1)
	go listApiProducts(&wg)
	wg.Add(1)
	go getMps(&wg)
	wg.Wait()
}
