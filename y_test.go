package main

import (
	"testing"
)

func TestJsonToYaml(t *testing.T) {
	ja := []byte(`{ "x": {"blah": 2, "a": [{ "hello": "bonjour"}, 2, 3] }}`)
	ya := `---
x:
  a:
  - hello: bonjour
  - 2
  - 3
  blah: 2
`
	if JsonToYaml(ja) != ya {
		t.Errorf("Failed translating JSON object to YAML")
	}

	jb := []byte(`[{"birb": 1}]`)
	yb := "---\n- birb: 1\n"

	if JsonToYaml(jb) != yb {
		t.Errorf("Failed translating array of JSON to YAML")
	}
}
