package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
)

func main() {
	svc := ec2.New(session.New(), &aws.Config{Region: aws.String("sa-east-1")})
	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Tag Name", "Instance Id", "Instance Type", "AZ", "IP", "Status"})

	for _, res := range resp.Reservations {
		for _, instance := range res.Instances {
			var tag_name string
			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					tag_name = *tag.Value
				}
			}
			table.Append([]string{
				tag_name,
				*instance.InstanceId,
				*instance.InstanceType,
				*instance.Placement.AvailabilityZone,
				*instance.PrivateIpAddress,
				*instance.State.Name,
			})

		}
	}
	table.Render()

}
