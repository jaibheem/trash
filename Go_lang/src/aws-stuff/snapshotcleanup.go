package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	aws_access_key_id := os.Getenv("ACCESS_KEY_ID")
	aws_secret_access_key := os.Getenv("SECRET_ACCESS_KEY")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)
	svc := ec2.New(session.New(), &aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: creds,
	})
	fmt.Println(reflect.TypeOf(svc))
	regions, _ := svc.DescribeRegions(nil)
	for _, region := range regions.Regions {
		regionname := *region.RegionName
		svc := ec2.New(session.New(), &aws.Config{
			Region:      aws.String(regionname),
			Credentials: creds,
		})
		fmt.Println("Region:", regionname)
		params := &ec2.DescribeSnapshotsInput{
			DryRun: aws.Bool(false),
			Filters: []*ec2.Filter{
				{
					Name: aws.String("owner-id"),
					Values: []*string{
						aws.String("123456789"),
					},
				},
			},
		}

		resp, err := svc.DescribeSnapshots(params)
		if err != nil {
			fmt.Println("Describe snapshot error", err)
		}
		for _, snap := range resp.Snapshots {
			delete := true
			stime := snap.StartTime
			ctime := time.Now()
			diff := int((ctime.Sub(*stime).Hours()) / 24.0)
			if !(strings.Contains(*snap.Description, "Created by CreateImage") || strings.Contains(*snap.Description, "Copied for DestinationAmi")) {
				if diff > 10 {
					for _, tag := range snap.Tags {
						if *tag.Key == "flags" {
							tagvalue := *tag.Value
							if strings.Contains(tagvalue, "preserve=true") {
								delete = false
							}
						}
					}
				} else {
					delete = false
				}
				if delete {
					fmt.Println(*snap.SnapshotId, diff)
					snapparams := &ec2.DeleteSnapshotInput{
						DryRun:     aws.Bool(false),
						SnapshotId: aws.String(*snap.SnapshotId),
					}
					_, result := svc.DeleteSnapshot(snapparams)
					if result != nil {
						fmt.Println("failed to delete the snapshot with", result)
					}
				}
			} //else {
			//				fmt.Println(*snap.SnapshotId, diff)
			//}
		}
	}
}
