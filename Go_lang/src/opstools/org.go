package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
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

func listEnv() (environments []string) {
	envs, _ := getRequests("/o/jai/e")
	json.Unmarshal(envs, &environments)
	return environments
}

func listApiProducts() (apiproducts []string) {
	apis, _ := getRequests("/o/jai/apis")
	json.Unmarshal(apis, &apiproducts)
	return apiproducts
}

func getUsers() (userroles []string) {
	roles, _ := getRequests("/o/jai/userroles")
	json.Unmarshal(roles, &userroles)
	return userroles
}

func getPods() (p []Pods) {
	regions, _ := getRequests("/o/jai/pods")
	json.Unmarshal(regions, &p)
	return p
}

func getServers() (mps []string) {
	//      var mps []string
	//var e Expand
	mpUUIDs, _ := getRequests("/o/jai/e/prod/servers")
	json.Unmarshal(mpUUIDs, &mps)
	return mps
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
	uuids := getServers()
	for _, uuid := range uuids {
		var e Expand
		server, _ := getRequests("/servers/" + uuid)
		json.Unmarshal(server, &e)
		fmt.Println(e.ExternalHostName, e.ExternalIp, e.InternalHostName, e.InternalIp)
	}
	//Print userroles
	//	userroles := getUsers()
	//	for _, i := range userroles {
	//		var allroles []string
	//		role, _ := getRequests("/o/jai/userroles/" + i + "/users")
	//		json.Unmarshal(role, &allroles)
	//		fmt.Printf("[" + i + "]")
	//		fmt.Printf("\t")
	//		for _, j := range allroles {
	//			fmt.Printf(j)
	//			fmt.Printf("\t")
	//		}
	//		fmt.Println()
	//	}
	//	//Print API Products
	//	apiproducts := listApiProducts()
	//	for _, api := range apiproducts {
	//		fmt.Printf(api)
	//		fmt.Printf("\t")
	//	}
	//	//Print Envs
	//	envs := listEnv()
	//	for _, env := range envs {
	//		fmt.Printf(env)
	//		fmt.Printf("\t")
	//	}
	//	//Print Pod and Region
	//	p := getPods()
	//	for _, i := range p {
	//		fmt.Println("\n", i.Name, "\t", i.Region)
	//	}
}
