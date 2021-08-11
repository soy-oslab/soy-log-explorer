package global

import (
	"testing"
)

func TestInit(t *testing.T) {
	if ColdRing == nil {
		t.Error("No ColdRing")
	}
}
