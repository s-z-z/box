package cprt

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

func PrettyJson(o interface{}) string {
	jsonBytes, _ := json.MarshalIndent(o, "", "  ")
	return string(jsonBytes)
}

func Yaml(o interface{}) string {
	bytes, _ := yaml.Marshal(o)
	return string(bytes)
}

func YamlInfo(o interface{}) {
	Info(Yaml(o))
}

func YamlWarning(o interface{}) {
	Warning(Yaml(o))
}
func YamlDebug(o interface{}) {
	Debug(Yaml(o))
}

func YamlError(o interface{}) {
	Error(Yaml(o))
}
