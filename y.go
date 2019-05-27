package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	bytes, err := ioutil.ReadAll(os.Stdin)
	check(err)

	var parsedJson map[string]interface{}
	json.Unmarshal([]byte(string(bytes)), &parsedJson)
	yamlOutput, err := yaml.Marshal(&parsedJson)
	check(err)

	fmt.Println(string(yamlOutput))
}
