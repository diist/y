package main

import (
	"testing"
)

func TestJsonToYaml(t *testing.T) {
	j := []byte(`{ "x": {"blah": 2, "a": [{ "hello": "bonjour"}, 2, 3] }}`)
	y := `x:
  a:
  - hello: bonjour
  - 2
  - 3
  blah: 2
`
	if JsonToYaml(j) != y {
		t.Errorf("Failed translating JSON to YAML")
	}
}
