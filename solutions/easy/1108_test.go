package easy

import (
	"fmt"
	"testing"
)

func TestDefangIPaddr(t *testing.T) {
	address := "192.168.0.1"
	fmt.Println(defangIPaddr(address))
}
