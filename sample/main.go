package sample

import (
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

func main() {
	_, err := jsonpatch.DecodePatch(nil)
	if err != nil {
		fmt.Println(err)
	}
}
