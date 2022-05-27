package main

import (
	"os/exec"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

func main() {
	action := githubactions.New()

	releaseName := action.GetInput("release-name")
	if releaseName == "" {
		action.Fatalf("release-name is required")
	}

	chartPath := action.GetInput("chart-path")
	if chartPath == "" {
		action.Fatalf("chart-path is required")
	}

	action.Infof("Running: %s")

	args := []string{"upgrade", "--install", releaseName, chartPath}
	cmd := exec.Command("helm", args...)

	cmdString := strings.Join(args, " ")
	action.Infof("Running: helm %s", cmdString)

	if err := cmd.Run(); err != nil {
		action.Fatalf("helm upgrade failed: %s", err)
	}
}
