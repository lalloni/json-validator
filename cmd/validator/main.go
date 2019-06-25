package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pkg/errors"

	validator "gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator.git"
	"gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator.git/schemas"
)

func main() {
	schema := "persona"
	flag.StringVar(&schema, "schema", "persona", "The schema to validate with (one of "+strings.Join(schemas.List(), ", ")+")")
	flag.Parse()
	dec := json.NewDecoder(os.Stdin)
	for {
		data, err := next(dec)
		if err == io.EOF {
			return
		}
		if err != nil {
			fatal("reading next document: %v", err)
		}
		report, err := validate(data, schema)
		if err != nil {
			fatal("validating document: %v", err)
		}
		err = print(report)
		if err != nil {
			fatal("printing validation report: %v", err)
		}
	}
}

func next(dec *json.Decoder) (map[string]interface{}, error) {
	data := map[string]interface{}{}
	err := dec.Decode(&data)
	return data, err
}

func validate(data map[string]interface{}, schema string) (map[string]interface{}, error) {
	report := make(map[string]interface{})
	if id, ok := data["id"]; ok {
		report["id"] = id
	}
	bs, err := json.Marshal(&data)
	if err != nil {
		return report, errors.Wrapf(err, "error encoding next: %v", err)
	}
	res, err := validator.ValidateJSON(schemas.MustLoad(schema), bs)
	if err != nil {
		return report, errors.Wrapf(err, "error validating json document: %v", err)
	}
	report["valid"] = res.Valid()
	if !res.Valid() {
		report["errors"] = res.Errors
	}
	return report, nil
}

func print(r map[string]interface{}) error {
	bs, err := json.Marshal(r)
	if err != nil {
		return errors.Wrap(err, "marshalling report for printing")
	}
	_, err = fmt.Println(string(bs))
	return err
}

func fatal(f string, arg ...interface{}) {
	f = "fatal: " + f
	if !strings.HasSuffix(f, "\n") {
		f += "\n"
	}
	fmt.Fprintf(os.Stderr, f, arg...)
	os.Exit(1)
}
