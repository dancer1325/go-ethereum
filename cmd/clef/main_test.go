package main

import (
	"testing"

	"github.com/urfave/cli/v2"
)

func TestInitializeSecrets(t *testing.T) {
	flagSet, _ := cli.flagSet("stdio-ui")
	parentContext := cli.NewContext(nil, nil, nil)
	context := cli.NewContext(nil, flagSet, parentContext)
	err := initializeSecrets(context)
	if err != nil {
		return
	}
}

func ExampleDefaultConfigDir() {
	DefaultConfigDir()
}
