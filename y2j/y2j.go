package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	err := y2j()
	if err != nil {
		log.Fatal(err)
	}
}

func y2j() error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	b, err = yaml.YAMLToJSON(b)
	if err != nil {
		return err
	}

	fmt.Print(string(b))
	return nil
}
