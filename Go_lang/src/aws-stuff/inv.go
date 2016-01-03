package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

var (
	wg    sync.WaitGroup
	table = tablewriter.NewWriter(os.Stdout)
)

func main() {
	region := []string{"ap-southeast-2", "sa-east-1", "ap-southeast-1", "ap-northeast-1", "eu-west-1", "us-west-2", "us-east-1"}
	table.SetHeader([]string{"Region", "c3.large", "c3.xlarge", "c3.2xlarge", "c3.4xlarge", "c3.8xlarge", "m1.large"})

	for _, reg := range region {
		wg.Add(1)
		go countInstances(reg)
	}
	wg.Wait()
	table.Render()
}

func countInstances(reg string) {
	//	aws_access_key_id := os.Getenv("aws_access_key")
	//	aws_secret_access_key := os.Getenv("aws_secret_key")
	aws_access_key_id := os.Getenv("ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("SECRET_ACCESS_KEY")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	svc := ec2.New(session.New(), &aws.Config{
		Region:      aws.String(reg),
		Credentials: creds,
	})
	instanceset := map[string]int{
		"c3.large":   0,
		"c3.xlarge":  0,
		"c3.2xlarge": 0,
		"c3.4xlarge": 0,
		"c3.8xlarge": 0,
		"m1.large":   0,
	}
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag-value"),
				Values: []*string{
					aws.String("cassandra"),
					aws.String("usergrid_cassandra"),
				},
			},
		},
	}
	resp, err := svc.DescribeInstances(params)
	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		panic(err)
		fmt.Println(err)
	}
	for _, reservation := range resp.Reservations {
		if len(reservation.Instances) > 0 {
			for _, instance := range reservation.Instances {
				instanceset[*instance.InstanceType]++
			}
		}
	}
	table.Append([]string{
		reg,
		strconv.Itoa(instanceset["c3.large"]),
		strconv.Itoa(instanceset["c3.xlarge"]),
		strconv.Itoa(instanceset["c3.2xlarge"]),
		strconv.Itoa(instanceset["c3.4xlarge"]),
		strconv.Itoa(instanceset["c3.8xlarge"]),
		strconv.Itoa(instanceset["m1.large"]),
	})

	wg.Done()
}
