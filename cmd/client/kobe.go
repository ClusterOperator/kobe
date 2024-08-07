package main

import (
	"github.com/ClusterOperator/kobe/cmd/client/root"
	"github.com/ClusterOperator/kobe/pkg/config"
	"os"
)

func main() {
	config.Init()
	if err := root.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
