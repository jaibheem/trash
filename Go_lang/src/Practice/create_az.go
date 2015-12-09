package main
import (
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/elb"
	"fmt"
	"os"
)
var region string = "eu-central-1"
type CreateLoadBalancer struct {
	Name              string
	AvailabilityZones []string
	Listeners         []Listener
	Scheme            string
	SecurityGroups    []string
	Subnets           []string
}
func main() {
	auth := aws.Auth{
		AccessKey: os.Getenv("ACCESS_KEY_ID"),
		SecretKey: os.Getenv("SECRET_ACCESS_KEY"),
	}
	e := elb.New(auth, aws.Region{ELBEndpoint: "https://elasticloadbalancing.amazonaws.com"})
	fmt.Println(elbName)
	resp, err := e.CreateLoadBalancer("testing")
	if err != nil {
		panic(err)
	}
	fmt.Printf(resp)
}