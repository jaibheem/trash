package main
import (
	"code.google.com/p/go.crypto/ssh"
	"fmt"
	//"os"
	"io/ioutil"
	"os/user"
	"bytes"
)
func getKeyFile() (key ssh.Signer, err error) {
	usr, _ := user.Current()
	file := usr.HomeDir + "/.ssh/jai_rsa"
	buf, err := ioutil.ReadFile(file)
	if err != nil{
		panic(err)
	}
	key, err = ssh.ParsePrivateKey(buf)
	if err != nil{
		panic(err)
	}
	return
}

func main() {
	key, err := getKeyFile()
	if err !=nil {
		panic(err)
	}
	config := &ssh.ClientConfig{
		User : "jaibheemsen",
		Auth: []ssh.AuthMethod{ssh.PublicKeys(key)},
	}
	client, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil{
		panic(err)
	}
	session, err := client.NewSession()
	if err != nil {
    	panic("Failed to create session: " + err.Error())
	}
	defer session.Close()
	var b bytes.Buffer
	session.Stdout = &b
	if err := session.Run("/usr/bin/whoami"); err != nil {
		panic(err.Error())
	}
	fmt.Println(b.String())
}//godoc -http=:1989 -index=true
