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

	other, err := row.ParseBlink()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(other)

	fmt.Println(row.Presets)
}
