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

func JsonToYaml(bytes []byte) string {
	var j interface{}
	json.Unmarshal(bytes, &j)

	y, e := yaml.Marshal(&j)
	check(e)

	return string(y)
}

func main() {
	bytes, e := ioutil.ReadAll(os.Stdin)
	check(e)

	s := JsonToYaml(bytes)
	fmt.Println(s)
}
