package main

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/elb"
)

var elbInfo = make(map[string]map[string]map[string]string)

func getElbInfo(elbName string, cli **elb.ELB) (*elb.DescribeLoadBalancersResult, error) {
	elbs := elb.DescribeAccessPointsInput{LoadBalancerNames: []string{elbName}}
	resp, err := cli.DescribeLoadBalancers(&elbs)
	return resp, err
}

func getInstanceHealth(InstanceIds []elb.Instance, cli **elb.ELB) { //(*elb.DescribeInstanceHealthResult, error) {
	for _, instance := range InstanceIds {
		id := elb.Instance{}
		id.InstanceID = instance.InstanceID
		instances := elb.DescribeEndPointStateInput{Instances: []elb.Instance{id}}
		resp, _ := cli.DescribeInstanceHealth(&instances)
		fmt.Println(*instance.InstanceID, resp)
	}
	//  return resp, err
}

func main() {
	creds := aws.Creds("ACCESS", "SECRET", "")
	cli := elb.New(creds, "us-east-1", nil)
	//  elbs := elb.DescribeAccessPointsInput{LoadBalancerNames: []string{"grinder-perf", "etcd-test"}}
	//  elbs.LoadBalancerNames = []string{"test"}
	//  resp, err := cli.DescribeLoadBalancers(&elbs)
	resp, err := getElbInfo("grinder-perf", &cli)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range resp.LoadBalancerDescriptions {
		//      elbInfo[*v.DNSName]
		getInstanceHealth(v.Instances, &cli)
		//      for _, instance := range v.Instances {
		//          health = getInstanceHealth(id, &cli)
		//          elbInfo[*v.DNSName][*instance.InstanceID]
		//          fmt.Println(elbInfo)
		//  }
		//      elbInfo["test"]["instance"] = string(*v.Instances[0].InstanceID)
		//  fmt.Println(*v.Instances[0].InstanceID)
		//      fmt.Println(v.AvailabilityZones)
	}
	fmt.Println("######################", elbInfo)
}
