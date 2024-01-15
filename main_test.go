package main

import "testing"

func TestMain(t *testing.T) {
	var actual = SimpleFunc()
	if SimpleFunc() != "OK" {
		t.Errorf("out put not expected. expectd=%s actual=%s", "OK", actual)
	}
}
