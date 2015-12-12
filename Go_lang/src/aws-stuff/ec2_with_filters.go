package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func main() {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("sa-east-1")})
	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("tag-value"),
				Values: []*string{
					aws.String("cassandra"),
				},
			},
		},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Hostname", "PrivateIpAddress", "InstanceType", "AvailabilityZone"})
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
	table.Render()
}
