package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	validator "gitlab.cloudint.afip.gob.ar/blockchain-team/padfed-validator"
)

func main() {
	v, err := validator.New()
	if err != nil {
		log.Fatalf("error creating validator: %v", err)
	}
	data := make(map[string]interface{})
	dec := json.NewDecoder(os.Stdin)
	for {
		err = dec.Decode(&data)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error decoding next: %v", err)
		}
		bs, err := json.Marshal(&data)
		if err != nil {
			log.Fatalf("error encoding next: %v", err)
		}
		res, err := v.ValidatePersonaJSON(bs)
		if err != nil {
			log.Fatalf("error validating json document: %v", err)
		}
		if res.Valid() {
			fmt.Println("valid")
		}
		for _, problem := range res.Errors {
			fmt.Printf("%s: %s\n", problem.Field, problem.Description)
		}
	}
}
