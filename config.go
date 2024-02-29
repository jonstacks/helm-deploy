package main

import (
	"errors"
	"strings"

	"github.com/sethvargo/go-githubactions"
)

func parseBool(action *githubactions.Action, name string) bool {
	return action.GetInput(name) == "true"
}

func parseStringSlice(action *githubactions.Action, name, delimiter string) []string {
	return strings.Split(action.GetInput(name), delimiter)
}

type Config struct {
	// Required
	ReleaseName string
	Chart       string

	// Optional
	Atomic           bool
	CleanupOnFail    bool
	CreateNamespace  bool
	Debug            bool
	DependencyUpdate bool
	DryRun           bool
	Force            bool
	KubeContext      string
	Namespace        string
	Sets             []string
	Timeout          string
	Values           []string
	Wait             bool
}

func NewFromInputs(action *githubactions.Action) (Config, error) {
	c := &Config{}

	c.Atomic = parseBool(action, "atomic")
	c.CleanupOnFail = parseBool(action, "cleanup-on-fail")
	c.CreateNamespace = parseBool(action, "create-namespace")
	c.Debug = parseBool(action, "debug")
	c.DependencyUpdate = parseBool(action, "dependency-update")
	c.DryRun = parseBool(action, "dry-run")
	c.Wait = parseBool(action, "wait")
	c.Force = parseBool(action, "force")
	c.KubeContext = action.GetInput("kube-context")
	c.Namespace = action.GetInput("namespace")
	c.Sets = parseStringSlice(action, "sets", " ")
	c.Timeout = action.GetInput("timeout")
	c.Values = parseStringSlice(action, "values", ",")

	c.ReleaseName = action.GetInput("release-name")
	if c.ReleaseName == "" {
		return *c, errors.New("release-name is required")
	}

	c.Chart = action.GetInput("chart")
	if c.Chart == "" {
		return *c, errors.New("chart is required")
	}

	return *c, nil
}

func (c Config) ToArgs() []string {
	args := []string{"upgrade", "--install"}

	if c.Atomic {
		args = append(args, "--atomic")
	}

	if c.CleanupOnFail {
		args = append(args, "--cleanup-on-fail")
	}

	if c.CreateNamespace {
		args = append(args, "--create-namespace")
	}

	if c.Debug {
		args = append(args, "--debug")
	}

	if c.DependencyUpdate {
		args = append(args, "--dependency-update")
	}

	if c.DryRun {
		args = append(args, "--dry-run")
	}

	if c.Wait {
		args = append(args, "--wait")
	}

	if c.Force {
		args = append(args, "--force")
	}

	if c.KubeContext != "" {
		args = append(args, "--kube-context", c.KubeContext)
	}

	if c.Namespace != "" {
		args = append(args, "--namespace", c.Namespace)
	}

	if c.Values != nil {
		for _, value := range c.Values {
			if value != "" {
				args = append(args, "--values", strings.TrimSpace(value))
			}
		}
	}

	if c.Sets != nil {
		for _, set := range c.Sets {
			if set != "" {
				args = append(args, "--set", strings.TrimSpace(set))
			}
		}
	}

	if c.Timeout != "" {
		args = append(args, "--timeout", c.Timeout)
	}

	args = append(args, c.ReleaseName)
	args = append(args, c.Chart)

	return args
}
