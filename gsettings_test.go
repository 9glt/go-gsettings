package gsettings

import (
	"testing"
)

func TestGet(t *testing.T) {
	value, err := Get("key")
	if err == nil {
		t.Fatal()
	}
	if value != "" {
		t.Fatal()
	}
	Set("key", "value")
	value, err = Get("key")
	if err != nil {
		t.Fatal()
	}
	if value != "value" {
		t.Fatal()
	}
}

func TestSet(t *testing.T) {
	Set("key", "value")
	v, exists := s["key"]
	if !exists {
		t.Fatal()
	}
	if v != "value" {
		t.Fatal()
	}
}
