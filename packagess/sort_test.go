package packagess

import (
	"fmt"
	"sort"
	"testing"
)


type user struct {
	Name string
	Age  int
}

type userSlice []user

func (uS userSlice) Len() int {
	return len(uS)
}

func (uS userSlice) Less(i, j int) bool {
	return uS[i].Age < uS[j].Age
}

func (uS userSlice) Swap(i, j int) {
	uS[i], uS[j] = uS[j], uS[i]
}
func TestSort(t *testing.T) {
	users := []user{
		{"Jalal", 27},
		{"Muh", 37},
		{"Akbar", 17},
	}

	fmt.Println("Before", users)

	sort.Sort(userSlice(users))

	fmt.Println("After", users)
}
