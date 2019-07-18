package gsettings_test

import (
	"testing"

	gsettings "github.com/9glt/go-gsettings"
)

func TestExternal(t *testing.T) {
	gsettings.Reset()
	value, err := gsettings.Get("key")
	if err == nil {
		t.Fatal()
	}
	if value != "" {
		t.Fatal()
	}
	gsettings.Set("key", "value")
	value, err = gsettings.Get("key")
	if err != nil {
		t.Fatal()
	}
	if value != "value" {
		t.Fatal()
	}
}
