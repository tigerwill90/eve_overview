package overviewsdk

import "testing"

func TestGetItems(t *testing.T) {
	if err := GetItems(); err != nil {
		t.Fatal(err)
	}
}
