package cloudfunc

import (
	"testing"
)

func TestA(t *testing.T) {

	if message != "Hello world" {
		t.Errorf("Message (%s) should be 'Hello world'", message)
	}
}
