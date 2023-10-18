//go:build mage

package main

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"

	"github.com/golangci/golangci-lint/pkg/commands"
	"github.com/magefile/mage/mg"
	"golang.org/x/vuln/scan"
)

func init() {
	os.Setenv(mg.VerboseEnv, strconv.FormatBool(true))
}

func LintVerbose(ctx context.Context) error {
	oldArgs := make([]string, len(os.Args))
	copy(oldArgs, os.Args)
	os.Args = append([]string{"golangci-lint"}, "run", "-v", "./...")
	defer func() { os.Args = oldArgs }()

	info := commands.BuildInfo{}
	if buildInfo, available := debug.ReadBuildInfo(); available {
		info.GoVersion = buildInfo.GoVersion
		info.Version = buildInfo.Main.Version
		info.Commit = fmt.Sprintf("(unknown, mod sum: %q)", buildInfo.Main.Sum)
		info.Date = "(unknown)"
	}

	return commands.NewExecutor(info).Execute()
}

func Lint(ctx context.Context) error {
	oldArgs := make([]string, len(os.Args))
	copy(oldArgs, os.Args)
	os.Args = append([]string{"golangci-lint"}, "run", "./...")
	defer func() { os.Args = oldArgs }()

	info := commands.BuildInfo{}
	if buildInfo, available := debug.ReadBuildInfo(); available {
		info.GoVersion = buildInfo.GoVersion
		info.Version = buildInfo.Main.Version
		info.Commit = fmt.Sprintf("(unknown, mod sum: %q)", buildInfo.Main.Sum)
		info.Date = "(unknown)"
	}

	return commands.NewExecutor(info).Execute()
}

// VulnCheck runs golang.org/x/vuln/scan with given args.
func VulnCheck(ctx context.Context) error {
	cmd := scan.Command(ctx, "./...")
	err := cmd.Start()
	if err == nil {
		err = cmd.Wait()
	}
	return err
}

func CheckAllVerboseLint(ctx context.Context) {
	mg.SerialCtxDeps(ctx, LintVerbose, VulnCheck)
}

func CheckAll(ctx context.Context) {
	mg.SerialCtxDeps(ctx, Lint, VulnCheck)
}
