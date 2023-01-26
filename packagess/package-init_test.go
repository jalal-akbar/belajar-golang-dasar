package packagess

import (
	"fmt"
	"testing"
)

var connection string

func init() {
	fmt.Println(" package init")
	connection = "Memanggil Fungsi init"
}

func GetInit() string {
	return connection
}

func TestPackageInit(t *testing.T) {
	result := GetInit()
	fmt.Println(result)
}
