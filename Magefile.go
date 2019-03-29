// +build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator/build"
)

func GenerateSchemas() error {
	return build.Convert("./schemas/", "./schemas/")
}

func GeneratePackr() error {
	return sh.Run("packr2", "-v", "--legacy")
}

func GenerateAll() error {
	mg.Deps(GenerateSchemas, GeneratePackr, GenerateVendor)
	return nil
}

func Test() error {
	mg.Deps(GenerateAll)
	return sh.RunV("go", "test", "-mod=vendor", "./...")
}

func Clean() error {
	return sh.Rm("target")
}

func Prepare() error {
	os.Mkdir("target", os.ModePerm)
	return nil
}

func GenerateVendor() error {
	return sh.Run("go", "mod", "vendor")
}

func Check() error {
	return sh.RunV("golangci-lint", "run")
}

func Compile() error {
	mg.Deps(GenerateAll)
	return sh.Run("go", "build", "-mod=vendor")
}

func Convey() error {
	return sh.RunV("goconvey", "-port=9999", "-watchedSuffixes=.go,.yaml", "-packages=1")
}

func Release() error {
	return errors.New("not implemented")
}
