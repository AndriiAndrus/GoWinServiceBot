package extra

import (
	"fmt"
	"testing"
)

func TestBypass(t *testing.T) {
	x := BypassAV()
	if x == true {
		fmt.Println("Debugger Is Present")
	}
}
