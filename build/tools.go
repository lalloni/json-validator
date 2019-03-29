package build

import (
	"log"
	"os/exec"

	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
)

func RunTool(command string, args []string, env map[string]string, installer func() error) error {
	path, err := exec.LookPath(command)
	if err != nil {
		if installer == nil {
			return errors.Errorf("running %s: not found", command)
		}
		log.Printf("running installer for %s", command)
		err := installer()
		if err != nil {
			return err
		}
		path, err = exec.LookPath(command)
		if err != nil {
			return errors.Errorf("installing %s: not found after running installer", command)
		}
	}
	return sh.RunWith(env, path, args...)
}

func RunPackr(args ...string) error {
	env := map[string]string{"GO111MODULE": "on"}
	return RunTool("packr2", args, env, InstallPackr)
}

func InstallPackr() error {
	return sh.Run("go", "get", "github.com/gobuffalo/packr/v2/packr2")
}

func RunLinter(args ...string) error {
	return RunTool("golangci-lint", args, nil, InstallLinter)
}

func InstallLinter() error {
	return sh.Run("go", "get", "github.com/golangci/golangci-lint/cmd/golangci-lint")
}

func RunGoConvey(args ...string) error {
	return RunTool("goconvey", args, nil, InstallGoConvey)
}

func InstallGoConvey() error {
	return sh.Run("go", "get", "github.com/smartystreets/goconvey")
}
