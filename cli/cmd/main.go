package main

import (
	"os"
	"time"

	"github.com/solo-io/sqoop/version"

	check "github.com/solo-io/go-checkpoint"
	"github.com/solo-io/sqoop/cli/pkg/cmd"
)

func main() {
	start := time.Now()
	defer check.CallCheck("sqoopctl", version.Version, start)

	app := cmd.App(version.Version)
	if err := app.Execute(); err != nil {
		if err != nil {
			os.Exit(1)
		}
		os.Exit(0)
	}
}
