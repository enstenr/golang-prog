package main

import (
	"github.com/bmuschko/kubectl-server-version/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

var version = "undefined"

func main() {
	cmd.SetVersion(version)

	serverVersionCmd := cmd.NewServerVersionCommand(genericclioptions.IOStreams{In: os.Stdin, Out: os.Stdout, ErrOut: os.Stderr})
	if err := serverVersionCmd.Execute(); err != nil {
		os.Exit(1)
	}
}