package cmd

import (
	"github.com/urfave/cli/v2"
	"k8s-playground/util/version"
)

var Cmds []*cli.Command

func init() {
	version := &cli.Command{
		Name:   "version",
		Usage:  "Displays K8S-Playground CLI version",
		Action: VersionLookup,
	}
	Cmds = append(Cmds, version)
}

func VersionLookup(_ *cli.Context) error {
	version.Lookup()
	return nil
}

