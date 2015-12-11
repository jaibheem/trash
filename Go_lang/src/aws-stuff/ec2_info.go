package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("sa-east-1")})
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}
	for _, res := range resp.Reservations {
		for _, instance := range res.Instances {
			fmt.Printf("%s --> %s\n", *instance.InstanceId, *instance.InstanceType)
		}
	}
}
