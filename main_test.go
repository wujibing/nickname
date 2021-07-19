package nickname

import "testing"

func TestGet(t *testing.T) {
	for i := 0; i < 100; i ++ {
		t.Log(Get())
	}
}