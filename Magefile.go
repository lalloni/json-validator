// +build mage

package main

import (
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator/build"
)

// Genera directorio vendor con las dependencias
func Genvendor() error {
	return sh.Run("go", "mod", "vendor")
}

// Genera JSON Schema
func Genschema() error {
	return build.Convert("schemas/resources", "doc/schemas")
}

// Genera recursos embebidos en código fuente
func Genpack() error {
	return build.RunPackr("-v")
}

// Genera todos los "generables"
func Genall() error {
	mg.Deps(Genschema, Genpack, Genvendor)
	return nil
}

// Ejecuta tests
func Test() error {
	mg.Deps(Genall)
	return sh.RunV("go", "test", "-mod=vendor", "./...")
}

// Ejecuta análisis estático de código fuente
func Check() error {
	return build.RunLinter("run")
}

// Ejecuta compilación de código fuente
func Compile() error {
	mg.Deps(Genall)
	return sh.Run("go", "build", "-mod=vendor")
}

// Lanza GoConvey (http://goconvey.co/)
func Convey() error {
	return build.RunGoConvey("-port=9999", "-watchedSuffixes=.go,.yaml", "-packages=1")
}

// Ejecuta el proceso de release
func Release() error {
	mg.Deps(Genall, Compile, Test, Check)
	return errors.New("not implemented")
}

// Construye un binario estático de este build
func Buildbuild() error {
	return sh.RunV("mage", "-compile", "magestatic")
}
