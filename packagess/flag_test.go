package packagess

import (
	"flag"
	"fmt"
	"testing"
)

func TestFlag(t *testing.T) {
	var host *string = flag.String("host", "localhost", "Put Your Host")
	var username *string = flag.String("username", "root", "Put Your Username")
	var password *string = flag.String("password", "root", "Put Your Password")

	flag.Parse()

	fmt.Println("Host :", *host)
	fmt.Println("Username :", *username)
	fmt.Println("Password :", *password)

}
