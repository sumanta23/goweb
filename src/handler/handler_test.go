package handler
import (
	"testing"
)

func assertEqual(t *testing.T, a interface{}, b interface{}) {
	if a != b {
		t.Fatalf("%s != %s", a, b)
	}
}

func TestAbc(t *testing.T) {
	  assertEqual(t, 1,1)
}
