package main

import "testing"

func TestHandleLifecycle(t *testing.T) {
	h := SpiceNukedOPL3_Create(44100)
	if h == 0 {
		t.Fatal("expected non-zero handle")
	}
	if rc := SpiceNukedOPL3_Destroy(h); rc != 0 {
		t.Fatalf("destroy returned %d", rc)
	}
}
