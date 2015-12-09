package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/elb"
)

func getAllElbs(elbs string, cli *elb.ELB) (*elb.DescribeLoadBalancersResult, error) {
	elbinput := elb.DescribeAccessPointsInput{LoadBalancerNames: []string{elbs}}
	resp, err := cli.DescribeLoadBalancers(&elbinput)
	return resp, err
}

func main() {
	creds := aws.Creds(os.Getenv("research_key"), os.Getenv("research_secret"), "")
	cli := elb.New(creds, "us-east-1", nil)
	resp, err := getAllElbs("grinder-perf", cli)
	fmt.Println(reflect.TypeOf(resp))
	if err != nil {
		fmt.Println(err)
	}
	for _, instance := range resp.LoadBalancerDescriptions {
		fmt.Println(*instance.LoadBalancerName)
		instanceids := instance.Instances
		fmt.Println(reflect.TypeOf(&instanceids))
	}
	//	fmt.Println(reflect.TypeOf(&instanceids))
	//	instanceinput := elb.DescribeEndPointStateInput{}
	//  elbs := elb.DescribeAccessPointsInput{LoadBalancerNames: []string{"grinder-perf", "etcd-test"}}

}
