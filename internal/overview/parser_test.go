package overview

import (
	"fmt"
	"os"
	"testing"
)

func TestRawUnmarshal(t *testing.T) {
	f, err := os.Open("../../all_color.yaml")
	if err != nil {
		t.Fatal(err)
	}
	row, err := RawUnmarshal(f)
	if err != nil {
		t.Fatal(err)
	}

	res, err := row.ParseColors()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)

	other, err := row.ParseBlinks()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(other)

	another, err := row.ParsePresets()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(another)
}
