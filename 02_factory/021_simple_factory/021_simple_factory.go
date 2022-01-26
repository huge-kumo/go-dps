package factory

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
)

// IRuleConfigParser 实例
type IRuleConfigParser interface {
	Parse(data []byte)
}

// jsonRuleConfigParser jsonRuleConfigParser
type jsonRuleConfigParser struct{}

func (J jsonRuleConfigParser) Parse(data []byte) {
	jsonMap := map[string]interface{}{}

	_ = json.Unmarshal(data, &jsonMap)
	fmt.Println(jsonMap)
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct{}

func (Y yamlRuleConfigParser) Parse(data []byte) {
	yamlMap := map[string]interface{}{}

	_ = yaml.Unmarshal(data, &yamlMap)
	fmt.Println(yamlMap)
}

// NewIRuleConfigParser 简单工厂
func NewIRuleConfigParser(t string) IRuleConfigParser {
	switch t {
	case "json":
		return jsonRuleConfigParser{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil
}
