package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	aws_access_key_id := os.Getenv("aws_access_key")
	aws_secret_access_key := os.Getenv("aws_secret_key")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	region := []string{"ap-southeast-2", "sa-east-1", "ap-southeast-1", "ap-northeast-1", "eu-west-1", "us-west-2", "us-east-1"}
	xl4 := 0
	xl2 := 0
	xl := 0
	other := 0
	for _, reg := range region {
		svc := ec2.New(session.New(), &aws.Config{
			Region:      aws.String(reg),
			Credentials: creds,
		})
		params := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name: aws.String("tag-value"),
					Values: []*string{
						aws.String("application2"),
						aws.String("application1"),
						aws.String("application3"),
					},
				},
			},
		}
		resp, err := svc.DescribeInstances(params)
		if err != nil {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
			return
		}
		for _, reservation := range resp.Reservations {
			if len(reservation.Instances) > 0 {
				for _, instance := range reservation.Instances {
					if *instance.InstanceType == "c3.2xlarge" {
						xl2 = xl2 + len(reservation.Instances)
					} else if *instance.InstanceType == "c3.4xlarge" {
						xl4 = xl4 + len(reservation.Instances)
					} else if *instance.InstanceType == "c3.xlarge" {
						xl = xl + len(reservation.Instances)
					} else {
						other = other + len(reservation.Instances)
					}
				}
			}
		}
	}
	if xl4 > 0 {
		fmt.Println("Total number of C3.4xlarge Nodes --> ", xl4)
	}
	if xl2 > 0 {
		fmt.Println("Total number of C3.2xlarge Nodes --> ", xl2)
	}
	if xl > 0 {
		fmt.Println("Total number of C3.xlarge Nodes --> ", xl)
	}
	if other > 0 {
		fmt.Println("Total number of other instancetype Nodes --> ", other)
	}

}
