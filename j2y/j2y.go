package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	err := j2y()
	if err != nil {
		log.Fatal(err)
	}
}

func j2y() error {
	b, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return err
	}

	b, err = yaml.JSONToYAML(b)
	if err != nil {
		return err
	}

	fmt.Print(string(b))
	return nil
}
