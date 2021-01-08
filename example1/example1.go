package stuff

import(
	"fmt"
	"strings"
)

func ex0() string {
	fmt.Printf("here")
	if strings.Contains("this is mouse", "mouse") {
		return "MOUSE"
	}
	return "no Mouse"
}