package facade

import (
	"testing"
)

var expect = "A module running\nB module running"

func TestNewAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != expect {
		t.Fatalf("except %s, return %s", expect, ret)
	}
}
