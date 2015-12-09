package main

import (
	"fmt"
	"log"
	"os"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/elb"
)

var regions = []string{

	"us-east-1",
	"us-west-2",
	"eu-west-1",
	"ap-southeast-1",
	"ap-northeast-1",
}

func getAllElbs(region string, auth *aws.Auth) ([]elb.LoadBalancerDescription, error) {
	endpoint := "https://elasticloadbalancing." + region + ".amazonaws.com"
	elbConn := elb.New(*auth, aws.Region{ELBEndpoint: endpoint})
	resp, err := elbConn.DescribeLoadBalancers()
	return resp.LoadBalancerDescriptions, err
}

func main() {
	auth := aws.Auth{
		AccessKey: os.Getenv("ACCESS_KEY_ID"),
		SecretKey: os.Getenv("SECRET_ACCESS_KEY"),
	}
	for _, region := range regions {
		elbs, err := getAllElbs(region, &auth)
		if err != nil {
			log.Print(err)
		}
	for _, elb := range elbs {
		fmt.Printf("%s\n", elb.LoadBalancerName)
	}
	}
}