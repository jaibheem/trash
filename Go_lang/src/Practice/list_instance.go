package main

import (
	"fmt"
	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/ec2"
	"os"
	"reflect"
)

type Tag struct {
	Name *string `locationName:"key" type:"string"`
}

func main() {
	AccessKey := os.Getenv("AWS_ACCESS_KEY")
	SecretKey := os.Getenv("AWS_SECRET_KEY")
	creds := aws.Creds(AccessKey, SecretKey, "")

	svc := ec2.New(&aws.Config{Credentials: creds, Region: "ap-southeast-1"})
	resp, err := svc.DescribeInstances(nil)
	fmt.Println(reflect.TypeOf(resp))
	if err != nil {
		panic(err)
	}

	//fmt.Println(resp.Reservations)
	for _, reservation := range resp.Reservations {
		//fmt.Println(reservation.Instances)
		for _, instance := range reservation.Instances {
			for _, t := range instance.Tags {
				if *t.Value == "cassandra" {
					fmt.Print(*instance.InstanceID, "\t", *instance.InstanceType, "\t", *instance.PrivateIPAddress, "\t")
					fmt.Println(*t.Value)
					fmt.Println()
				}
			}
		}
	}
}
