package main

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	wg           sync.WaitGroup
	cluster_name = os.Args[1]
)

func main() {
	region := []string{"ap-southeast-2", "sa-east-1", "ap-southeast-1", "ap-northeast-1", "eu-west-1", "us-west-2", "us-east-1"}
	for _, reg := range region {
		wg.Add(1)
		go getInstances(reg)
	}
	wg.Wait()
}

func getInstances(region string) {
	aws_access_key_id := os.Getenv("ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("SECRET_ACCESS_KEY")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	svc := ec2.New(session.New(), &aws.Config{
		Region:      aws.String(region),
		Credentials: creds,
	})
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag:cluster_name"),
				Values: []*string{
					aws.String(cluster_name),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		panic(err)
	}
	fmt.Println("[", region, "]")
	for _, reservation := range resp.Reservations {
		for _, instance := range reservation.Instances {
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					fmt.Println(strings.Split(*tag.Value, ".")[0])
				}
			}
		}
	}
	wg.Done()
}
