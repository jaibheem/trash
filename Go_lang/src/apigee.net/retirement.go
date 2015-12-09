package main

import (
	"fmt"
	"os"

	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/ec2"
)

var regions = []string{
	"us-east-1",
	"us-west-2",
	"eu-west-1",
	"ap-southeast-1",
	"ap-northeast-1",
}

func getInstanceStatus(region string, auth *aws.Auth, instances []string, filter *ec2.Filter) ([]ec2.InstanceStatus, error) {
	endpoint := "https://ec2." + region + ".amazonaws.com"
	ec2 := ec2.New(*auth, aws.Region{EC2Endpoint: endpoint})
	resp, err := ec2.DescribeInstanceStatus(nil, filter)
	return resp.InstanceStatuses, err
}

func main() {
	auth := aws.Auth{
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("AWS_SECRET_KEY"),
	}
	filter := ec2.NewFilter()
	filter.Add("event.code", "instance-retirement")
	for _, _region := range regions {
		region := _region
		statuses, err := getInstanceStatus(region, &auth, nil, filter)
		if err != nil {
			fmt.Println(err)
		}
		for _, status := range statuses {
			fmt.Printf("%s\t%s\t%s\t%s\t%s\n", status.InstanceId, status.AvailabilityZone, status.InstanceState.Name, status.SystemStatus, status.EventDetails)
		}
	}
}
