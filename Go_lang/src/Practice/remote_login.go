package main

import (
	"fmt"
	"os/user"
)
func currentUsername() *user.User{
	u, err := user.Current()
	if err != nil{
		panic("No such user" + err.Error())
	}
	return u
}
func main() {
	fmt.Printf("Current system user is: %s\n", currentUsername().Username)	
}//godoc -http=:1989 -index=true
