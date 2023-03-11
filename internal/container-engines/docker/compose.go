package docker

import (
	"gopkg.in/yaml.v2"
	"io"
)

type compose struct {
	services map[string]interface{}
}

func (c compose) ToYaml(writer io.Writer) error {
	encode := yaml.NewEncoder(writer)
	return encode.Encode(c)
}
