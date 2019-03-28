// +build mage

package main

import (
	"github.com/magefile/mage/mg"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator/build"
)

// Generate JSON schemas in ./schemas/ from YAML sources
func Generate() error {
	return build.Convert("./schemas/", "./schemas/")
}

func Test() error {
	mg.Deps(Generate)
	return nil
}
