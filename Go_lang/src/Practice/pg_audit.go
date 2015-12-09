package main

import (
	"fmt"
	"os"
	"bufio"
	"os/exec"
	"strings"
	"sync"
	"time"
)


func check(cmd string, hostname string) {
	out, err := exec.Command("sh","-c",cmd).Output()
	if( err != nil) {
		fmt.Println(err)
		fmt.Println(hostname)
	}else {
		if !strings.Contains(string(out), "java-1.7.0-openjdk-1.7.0.75") {
			fmt.Println(hostname)
		}else{
		}
	}
}
func main(){
        file, err := os.Open("hosts")
	if err != nil {
		panic(err)
	}
	defer file.Close()
        var wg sync.WaitGroup
	defer wg.Done()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
    		hostname := scanner.Text()
		cmd := "ansible all -i " + hostname + ", -m shell -a 'rpm -qa | grep java' -s"
		wg.Add(1)
		time.Sleep(2*time.Second)
		go check(cmd, hostname)
	}
	wg.Wait()
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}