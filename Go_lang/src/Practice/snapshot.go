package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/cheggaaa/pb"
	"github.com/crowdmob/goamz/aws"
	"github.com/crowdmob/goamz/ec2"
	"gopkg.in/alecthomas/kingpin.v1"
)

var (
	ec2_region   = kingpin.Flag("region", "ec2 region").String()
	awsAccessKey = kingpin.Flag("aws_access_key", "AWS access key. If not set then the value of the ACCESS_KEY_ID environment variable is used.").String()
	awsSecretKey = kingpin.Flag("aws_secret_key", "AWS secret key. If not set then the value of the SECRET_ACCESS_KEY environment variable is used.").String()
)

var (
	username string = "rraithatha@apigee.com"
	password string = config.Password
	host     string = "smtp.gmail.com"
	port     int    = 0
)

type snapshotConfig struct {
	SnapshotRetries int
	Password        string
	MailUser        string
	MailHost        string
	MailPort        int
}

var config snapshotConfig
var logFile = "/var/log/snapshots.log"
var t = time.Now().UTC().AddDate(0, 0, -15)
var currentTime = t.Format("2006-01-02T15:04:05.000Z")

type volumeSnapshotsType map[string][]map[string]map[string]string

var volumeSnapshots = make(volumeSnapshotsType)
var countSnapshots = make(map[string]int)

const maxConcurrency = 30

var throttle = make(chan int, maxConcurrency)

var ec *ec2.EC2

func getSnapshots(vol string, ids []string, wg *sync.WaitGroup) {
	defer wg.Done()
	filter := ec2.NewFilter()
	filter.Add("volume-id", vol)
	snapshots, err := ec.Snapshots(ids, filter)
	countSnapshots[vol] = len(snapshots.Snapshots)
	if err != nil {
		log.Println(err)
		return
	}
	for _, value := range snapshots.Snapshots {
		var listSnapshots = make(map[string]map[string]string)
		listSnapshots[value.Id] = make(map[string]string)
		listSnapshots[value.Id]["StartTime"] = value.StartTime
		volumeSnapshots[vol] = append(volumeSnapshots[vol], listSnapshots)
	}
}

func deleteOldSnapshots(volSnap *volumeSnapshotsType) {
	var wg sync.WaitGroup
	for _, v := range *volSnap {
		throttle <- 1
		wg.Add(1)
		for _, value := range v {
			for snapshot, snapshotTime := range value {
				if snapshotTime["StartTime"] < currentTime {
					go deleteSnapshot(snapshot, &wg)
				}
			}
		}
		<-throttle
	}
}

func createSnapshot(vol string, desc string, count int, wg *sync.WaitGroup) {
	if count > 0 {
		tryCount := config.SnapshotRetries - count + 1
		if wg != nil {
			defer wg.Done()
		}
		resp, err := ec.CreateSnapshot(vol, desc)
		if err != nil {
			log.Printf("Snapshot creation failed for %s - %s, %d try", vol, err, tryCount)
			count = count - 1
			time.Sleep(3 * time.Second)
			createSnapshot(vol, desc, count, nil)
			return
		}
		snap, _ := ec.Snapshots([]string{resp.Id}, nil)
		for snap.Snapshots[0].Status != "completed" {
			time.Sleep(3 * time.Second)
			snap, _ = ec.Snapshots([]string{resp.Id}, nil)
		}
		log.Printf("Snapshot %s created for volume id %s on try %d", snap.Snapshots[0].Id, vol, tryCount)
	}
}

func deleteSnapshot(ssid string, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := ec.DeleteSnapshots(ssid)
	if err != nil {
		log.Printf("Failed deleting %s - %s", ssid, err)
	} else {
		log.Printf("Deleted %s", ssid)
	}
}

func main() {
	if _, err := toml.DecodeFile("snapshot_backup.toml", &config); err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Println("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	kingpin.Version("0.0.1")
	kingpin.Parse()

	if *awsAccessKey == "" && *awsSecretKey == "" {
		*awsAccessKey = os.Getenv("ACCESS_KEY_ID")
		*awsSecretKey = os.Getenv("SECRET_ACCESS_KEY")
	} else if *awsAccessKey != "" && *awsSecretKey == "" || *awsAccessKey == "" && *awsSecretKey != "" {
		log.Println("Pass both aws_access_key and aws_secret_key or export ACCESS_KEY_ID and SECRET_ACCESS_KEY environment variables")
		os.Exit(1)
	}
	auth := aws.Auth{
		AccessKey: *awsAccessKey,
		SecretKey: *awsSecretKey,
	}

	region := aws.Region{
		EC2Endpoint: "https://ec2." + *ec2_region + ".amazonaws.com",
	}
	ec = ec2.New(auth, region)
	resp, _ := ec.DescribeVolumes(nil, nil)
	bar := pb.StartNew(len(resp.Volumes))
	for _, value := range resp.Volumes {
		bar.Increment()
		throttle <- 1
		wg.Add(1)
		wg2.Add(1)
		//		go createSnapshot(value.VolumeId, "Snapshot for"+value.VolumeId, 3, &wg)
		go createSnapshot("vol-f07e3e", "Snapshot for"+value.VolumeId, config.SnapshotRetries, &wg)
		go getSnapshots(value.VolumeId, nil, &wg2)
		<-throttle

	}
	wg.Wait()
	bar.FinishPrint("")
	//deleteOldSnapshots(&volumeSnapshots)

}//godoc -http=:1989 -index=true
