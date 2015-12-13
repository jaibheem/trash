package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func main() {
	aws_access_key_id := os.Getenv("aws_access_key")
	aws_secret_access_key := os.Getenv("aws_secret_key")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	region := []string{"ap-southeast-2", "sa-east-1", "ap-southeast-1", "ap-northeast-1", "eu-west-1", "us-west-2", "us-east-1"}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hostname", "PrivateIpAddress", "InstanceType", "AvailabilityZone"})
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
						aws.String("application1"),
						aws.String("application2"),
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
			for _, instance := range reservation.Instances {
				var tag_name string
				for _, tag := range instance.Tags {
					if *tag.Key == "Name" {
						tag_name = *tag.Value
					}

				}
				table.Append([]string{
					tag_name,
					*instance.PrivateIpAddress,
					*instance.InstanceType,
					*instance.Placement.AvailabilityZone,
				})
			}
		}
	}
	table.Render()

}
