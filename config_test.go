package main

import (
	"os"
	"strings"
	"testing"

	"github.com/c2fo/testify/assert"
	"github.com/sethvargo/go-githubactions"
)

func withEnv(t *testing.T, env map[string]string, f func()) {
	// Save the current environment
	preTest := map[string]string{}
	for _, e := range os.Environ() {
		kv := strings.Split(e, "=")
		preTest[kv[0]] = kv[1]
	}

	for k, v := range env {
		if err := os.Setenv(k, v); err != nil {
			t.Fatalf("failed to set env var %s: %s", k, err)
		}
	}

	f()

	for k := range env {
		if _, ok := preTest[k]; !ok {
			os.Unsetenv(k)
		} else {
			os.Setenv(k, preTest[k])
		}
	}
}

func TestValuesFiles(t *testing.T) {
	testEnv := map[string]string{
		"INPUT_VALUES":       "testdata/values.yaml, testdata/values2.yaml",
		"INPUT_CHART":        "mychart",
		"INPUT_NAMESPACE":    "mynamespace",
		"INPUT_RELEASE-NAME": "myrelease",
	}

	withEnv(t, testEnv, func() {
		config, err := NewFromInputs(githubactions.New())
		assert.NoError(t, err)

		args := config.ToArgs()
		assert.Equal(t, []string{"upgrade", "--install", "--namespace", "mynamespace", "--values", "testdata/values.yaml", "--values", "testdata/values2.yaml", "myrelease", "mychart"}, args)
	})
}
