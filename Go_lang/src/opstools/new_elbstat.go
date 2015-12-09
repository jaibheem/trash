package main

import (
	"fmt"
	"os"
	"reflect"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/gen/elb"
)

func descElb(elbname string, cli *elb.ELB) (*elb.DescribeLoadBalancersResult, error) {
	elbinput := elb.DescribeAccessPointsInput{LoadBalancerNames: []string{elbname}}
	resp, err := cli.DescribeLoadBalancers(&elbinput)
	return resp, err
}

func descInstanceHealth(instanceIds []elb.Instance, cli *elb.ELB) {
	for _, instance := range instanceIds {
		instanceid := elb.Instance{}
		instanceid.InstanceID = instance.InstanceID
		instances := elb.DescribeEndPointStateInput{Instances: []elb.Instance{instanceid}}
		resp, _ := cli.DescribeInstanceHealth(&instances)
		fmt.Println(*instance.InstanceID, resp.InstanceState)
		fmt.Println(reflect.TypeOf(resp))
	}
}

func main() {
	creds := aws.Creds(os.Getenv("research_key"), os.Getenv("research_secret"), "")
	cli := elb.New(creds, "us-east-1", nil)
	resp, err := descElb("grinder-perf", cli)
	if err != nil {
		fmt.Println(err)
	}
	//	fmt.Println(resp)
	for _, v := range resp.LoadBalancerDescriptions {
		descInstanceHealth(v.Instances, cli)
	}
	//fmt.Printf("%v", reflect.TypeOf(cli))
	//	args := os.Args[1:]
	//	for i := 0; i < len(args); i++ {
	//elbinput := elb.DescribeAccessPointsInput{LoadBalancerNames: []string{args[i]},}
	//elbs, err := cli.DescribeLoadBalancers(&elbinput)
	//elbs, err := cli.DescribeLoadBalancers(nil)
	//		if err != nil {
	//			fmt.Print(err)
	//		}
	//		for _, elb := range elbs.LoadBalancerDescriptions {
	//			fmt.Println("LoadBalancer Name:", *elb.LoadBalancerName)
	//			for _, instance := range elb.Instances {
	//				fmt.Println(*instance.InstanceID)
	//			}
	//		}
	//	}

}

func (c *ELB) DescribeInstanceHealth(req *DescribeEndPointStateInput) (resp *DescribeInstanceHealthResult, err error) {
resp = &DescribeInstanceHealthResult{}
err = c.client.Do("DescribeInstanceHealth", "POST", "/", req, resp)
return
}

func (c *ELB) DescribeLoadBalancers(req *DescribeAccessPointsInput) (resp *DescribeLoadBalancersResult, err error) {
resp = &DescribeLoadBalancersResult{}
err = c.client.Do("DescribeLoadBalancers", "POST", "/", req, resp)
return
}