package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

type YamlProblem struct {
	Difficulty string   `yaml:"difficulty"`
	Directory  string   `yaml:"directory"`
	Published  bool     `yaml:"published"`
	Date       *string  `yaml:"date,omitempty"`
	Tags       []string `yaml:"tags,omitempty"`
}

func ReadYamlProblems(path string) (map[string]YamlProblem, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return map[string]YamlProblem{}, err
	}

	problems := make(map[string]YamlProblem)
	err = yaml.Unmarshal(data, &problems)
	if err != nil {
		return map[string]YamlProblem{}, err
	}

	return problems, nil
}

func WriteYamlProblems(path string, problems map[string]YamlProblem) error {
	data, err := yaml.Marshal(problems)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}
