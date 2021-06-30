package global

import (
	"testing"
)

func TestInit(t *testing.T) {
	if HotRing == nil {
		t.Error("No HotRing")
	}
	if ColdRing == nil {
		t.Error("No ColdRing")
	}
}
