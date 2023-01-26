package packagess

import (
	"fmt"
	"os"
	"testing"
)

func TestOs(t *testing.T) {
	// Args hold the command-line arguments, starting with the program name.
	args := os.Args
	fmt.Println("Argumen : ", args)
	// Hostname returns the host name reported by the kernel.
	host, err := os.Hostname()
	if err != nil {
		fmt.Println("Error :", err.Error())
	} else {
		fmt.Println("Hostname: ", host)
	}
	// Getenv retrieves the value of the environment variable named by the key.
	// It returns the value, which will be empty if the variable is not present.
	// To distinguish between an empty value and an unset value, use LookupEnv.
	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")
	fmt.Println(username)
	fmt.Println(password)
}
