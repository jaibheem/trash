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
	region := []string{"ap-southeast-2", "sa-east-1", "ap-southeast-1", "ap-northeast-1", "eu-west-1", "us-west-2", "us-east-1"}
	for _, reg := range region {
		//		table := tablewriter.NewWriter(os.Stdout)
		//		table.SetHeader([]string{"c3.large", "c3.xlarge", "c3.2xlarge", "c3.4xlarge", "c3.8xlarge"})
		mapp := countInstances(reg)
		if len(mapp) != 0 {
			fmt.Println(mapp)
		}
	}
	//	region := []string{"us-west-2"}
}

func countInstances(reg string) map[string]int {
	//func countInstances(reg string) {
	aws_access_key_id := os.Getenv("aws_access_key")
	aws_secret_access_key := os.Getenv("aws_secret_key")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	instanceset := make(map[string]int)
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
					//			aws.String("application1"),
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
	//	fmt.Println(instanceset)
	return instanceset
}
