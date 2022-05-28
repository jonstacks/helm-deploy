package main

import (
	"os"
	"os/exec"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	action := githubactions.New()
	config, err := NewFromInputs(action)
	if err != nil {
		action.Fatalf("%v", err)
	}

	args := config.ToArgs()

	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	action.Infof("Running: %s", cmd.String())

	if err := cmd.Run(); err != nil {
		cmd.CombinedOutput()
		action.Fatalf("helm upgrade failed: %s", err)
	}
}
