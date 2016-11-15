package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ghodss/yaml"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "A YAML file is required")
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %s\n", err)
		os.Exit(2)
	}

	fmt.Fprintf(os.Stdout, "Original:\n%s\n", string(data))

	ydoc := map[string]interface{}{}
	if err := yaml.Unmarshal(data, &ydoc); err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing file: %s\n", err)
		os.Exit(3)
	}

	fmt.Printf("COFFEE: %q\n", ydoc["coffee"].(string))

	out, err := yaml.Marshal(ydoc)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error serializing file: %s\n", err)
		os.Exit(4)
	}

	fmt.Fprintln(os.Stdout, string(out))
}
