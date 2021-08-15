package overviewsdk

import (
	"fmt"
	"os"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	f, err := os.Open("all_color.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	ow, err := Unmarshal(f)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%+v\n", ow)
}
