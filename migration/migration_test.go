package migration

import "testing"

func TestNewVersion(t *testing.T) {
	v := NewVersion()
	if v == 0 {
		t.Error()
	}
}
