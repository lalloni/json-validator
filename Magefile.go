// +build mage

package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator.git/build"
)

func init() {
	log.Info("magefile initialized")
}

// Limpia directorio de proyecto de artefactos temporales generados
func Clean() {
	sh.Rm("target")
}

// Ejecuta tests
func Test() error {
	return sh.RunV("go", "test", "./...")
}

// Ejecuta análisis estático de código fuente
func Check() error {
	return build.RunLinter("run")
}

// Ejecuta análisis estático de código fuente y tests
func Verify() {
	mg.Deps(Check, Test)
}

// Ejecuta todas las tareas de compilcación
func Compile() error {
	return sh.Run("go", "build", "./...")
}

// Lanza GoConvey (http://goconvey.co/)
func Convey() error {
	err := build.RunPackr("clean")
	if err != nil {
		return err
	}
	return build.RunGoConvey("-port=9999", "-watchedSuffixes=.go,.yaml", "-packages=1")
}

// Ejecuta el proceso de release
func Release() error {
	log.Info("checking parameters")
	version := os.Getenv("ver")
	if version == "" {
		return errors.New(`Version is required for release.
You must set the version to be released using the environment variable 'ver'.
On unix-like shells you could do something like:
    env ver=1.2.3 mage release`)
	}
	if _, err := semver.NewVersion(version); err != nil {
		return errors.Wrapf(err, "checking syntax of version %q", version)
	}

	tag := "v" + version
	log.Infof("releasing version %s with tag %s", version, tag)

	log.Info("checking release tag does not exist")
	out, err := sh.Output("git", "tag")
	if err != nil {
		return errors.Wrap(err, "getting git tags")
	}
	s := bufio.NewScanner(strings.NewReader(out))
	for s.Scan() {
		if tag == s.Text() {
			return errors.Errorf("release tag %q already exists", tag)
		}
	}

	log.Info("updating generated resources")

	log.Info("checking working tree is not dirty")
	out, err = sh.Output("git", "status", "-s")
	if err != nil {
		return errors.Wrap(err, "getting git status")
	}
	if len(out) > 0 {
		return errors.Errorf("working directory is dirty")
	}

	log.Info("running linter, compiler & tests")
	mg.Deps(Compile, Check, Test)

	log.Infof("creating tag %s", tag)
	if err := sh.RunV("git", "tag", "-s", "-m", "Release "+version, tag); err != nil {
		return errors.Wrap(err, "creating git tag")
	}

	log.Infof("pushing tag %s to 'origin' remote", tag)
	if err := sh.RunV("git", "push", "origin", tag); err != nil {
		return errors.Wrap(err, "pushing tag to origin remote")
	}

	log.Infof("pushing current branch to 'origin' remote", tag)
	if err := sh.RunV("git", "push", "origin"); err != nil {
		return errors.Wrap(err, "pushing current branch to origin remote")
	}

	log.Info("release successfuly completed")

	return nil
}

// Construye un binario estático de este build
func Buildbuild() error {
	return sh.RunV("mage", "-compile", "magestatic")
}
