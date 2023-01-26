package packagess

import (
	"container/list"
	"fmt"
	"testing"
)

func TestContainerList(t *testing.T) {
	data := list.New()

	data.PushBack("Jalaluddin")
	data.PushBack("Muh")
	data.PushBack("Akbar")
	data.PushFront("Vania")

	fmt.Println("Front of Data : ", data.Front().Value)
	fmt.Println("Back of Data : ", data.Back().Value)
	fmt.Println("Next of Front : ", data.Front().Next().Value)
	fmt.Println("Previous of Back : ", data.Back().Prev().Value)
	// Forward
	for forward := data.Front(); forward != nil; forward = forward.Next() {
		fmt.Println("Forward : ", forward.Value)
	}
	// Reverse
	for reverse := data.Back(); reverse != nil; reverse = reverse.Prev() {
		fmt.Println("Reverse :", reverse.Value)
	}

}
