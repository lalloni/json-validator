// +build mage

package main

import (
	"fmt"
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator/build"
)

// Genera JSON Schema
func Genschema() error {
	return build.Convert("schemas/resources", "doc/schemas")
}

// Genera recursos embebidos en código fuente
func Genpack() error {
	return build.RunPackr()
}

// Genera todos los "generables"
func Genall() error {
	mg.Deps(Genschema, Genpack)
	return nil
}

// Ejecuta tests
func Test() error {
	mg.Deps(Genall)
	return sh.RunV("go", "test", "./...")
}

// Ejecuta análisis estático de código fuente
func Check() error {
	return build.RunLinter("run")
}

// Ejecuta compilación de código fuente
func Compile() error {
	mg.Deps(Genall)
	return sh.Run("go", "build", "./...")
}

// Lanza GoConvey (http://goconvey.co/)
func Convey() error {
	return build.RunGoConvey("-port=9999", "-watchedSuffixes=.go,.yaml", "-packages=1")
}

// Ejecuta el proceso de release
func Release() error {
	version := os.Getenv("version")
	if version == "" {
		return errors.New(`Version is required for release.
You must set the version to be released using the environment variable 'version'.
On unix-like shells you could do something like:
    env version=1.2.3 mage release`)
	}
	fmt.Printf("Releasing version: %s\n", version)
	mg.SerialDeps(Genall, Check, Compile, Test)
	return errors.New("still not implemented")
}

// Construye un binario estático de este build
func Buildbuild() error {
	return sh.RunV("mage", "-compile", "magestatic")
}
