package main

import (
	"os/exec"
	"testing"
)

func TestCLIAddList(t *testing.T) {
	cmd := exec.Command("go", "run", "./cmd/calculator", "add", "5+5")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("CLI add failed: %v", err)
	}

	cmd = exec.Command("go", "run", "./cmd/calculator", "list")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("CLI list failed: %v", err)
	}
	if len(out) == 0 {
		t.Fatal("expected some output from list")
	}
}
