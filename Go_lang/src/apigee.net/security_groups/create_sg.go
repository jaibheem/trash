package main

import (
	"fmt"
	"os"
	//"github.com/crowdmob/goamz/aws"
	//"github.com/crowdmob/goamz/ec2"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/ec2"
	// 	"reflect"
)

type SecurityGroup struct {
	Id          string `xml:"groupId"`
	Name        string `xml:"groupName"`
	Description string `xml:"groupDescription"`
	VpcId       string `xml:"vpcId"`
}

func main() {
	auth := aws.Auth{
		AccessKey: os.Getenv("research_key"),
		SecretKey: os.Getenv("research_secret"),
	}
	region := aws.Region{
		Name:        "us-east-1",
		EC2Endpoint: "https://ec2.us-east-1.amazonaws.com",
	}
	ec2conn := ec2.New(auth, region)
	//var sg SecurityGroup
	fmt.Printf("%+v\n", ec2conn)
	//	resp, _ := ec2conn.CreateSecurityGroup("testing12", "testing123")

	resp, err := ec2conn.CreateSecurityGroup(ec2.SecurityGroup{Name: "testing12",
		Description: "testing123", VpcId: "vpc-123456"})
	fmt.Printf("%+v\n", resp)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
}
