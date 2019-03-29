package schemas

import (
	"io/ioutil"
	"log"

	"github.com/gobuffalo/packr/v2"
	"github.com/lalloni/gojsonschema"
	"github.com/pkg/errors"

	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator/convert"
	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator/formats"
)

var fs = packr.New("schemas", "")

func init() {
	gojsonschema.Locale = locale{}
	gojsonschema.FormatCheckers.Add("cuit", formats.Cuit)
	gojsonschema.FormatCheckers.Add("periododiario", formats.PeriodoDiario)
	gojsonschema.FormatCheckers.Add("periodomensual", formats.PeriodoMensual)
	gojsonschema.FormatCheckers.Add("periodoanual", formats.PeriodoAnual)
}

func Persona() (*gojsonschema.Schema, error) {
	jsonloader, err := loaderFromYAML("persona.yaml")
	if err != nil {
		return nil, errors.Wrap(err, "building loader for persona json schema")
	}
	schemaloader := gojsonschema.NewSchemaLoader()
	schemaloader.Validate = true // validate schema
	schemaloader.Draft = gojsonschema.Draft7
	schema, err := schemaloader.Compile(jsonloader)
	if err != nil {
		return nil, errors.Wrap(err, "building json schema for persona")
	}
	schema.SetRootSchemaName("(persona)")
	return schema, nil
}

func loaderFromYAML(name string) (gojsonschema.JSONLoader, error) {
	f, err := fs.Open(name)
	if err != nil {
		return nil, errors.Wrapf(err, "opening '%s'", name)
	}
	bs, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrapf(err, "reading '%s'", name)
	}
	schema, err := convert.FromYAML(bs, convert.Options{Source: name})
	if err != nil {
		return nil, errors.Wrapf(err, "converting '%s' to JSON", name)
	}
	log.Printf("loaded json schema from '%s':\n%s", name, string(schema) )
	loader := gojsonschema.NewBytesLoader(schema)
	_, err = loader.LoadJSON() // for checking json
	if err != nil {
		return nil, errors.Wrapf(err, "parsing JSON converted from '%s'", name)
	}
	return loader, nil
}
