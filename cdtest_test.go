package cdtest

import "testing"

func TestAlwaysTrue(t *testing.T) {
	if alwaysTrue() == false {
		t.Fatal("alwaysTrue returned false")
	}
}
