package main

import (
	"os"
	"os/exec"

	"github.com/sethvargo/go-githubactions"
)

type Config struct {
	ReleaseName string
	Chart       string
}

func NewFromInputs(action *githubactions.Action) Config {

}

func main() {
	action := githubactions.New()

	releaseName := action.GetInput("release-name")
	if releaseName == "" {
		action.Fatalf("release-name is required")
	}

	chart := action.GetInput("chart")
	if chart == "" {
		action.Fatalf("chart is required")
	}

	args := []string{"upgrade", "--install", releaseName, chart}
	cmd := exec.Command("helm", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	action.Infof("Running: %s", cmd.String())

	if err := cmd.Run(); err != nil {
		cmd.CombinedOutput()
		action.Fatalf("helm upgrade failed: %s", err)
	}
}
