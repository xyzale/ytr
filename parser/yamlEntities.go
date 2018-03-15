package parser

var fileContent []byte

var yamlContent = Yaml{}

type Yaml struct {
	Commands map[string][]string   `yaml:"commands"`
}