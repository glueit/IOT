package main

import (
	"fmt"
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

type nextStateDetails struct {
	Name            string `yaml:"name"`
	NextTask        string `yaml:"nextTask"`
	NextTaskContext string `yaml:"nextTaskContext"`
}

type taskDetails struct {
	TaskName string           `yaml:"taskName"`
	State    nextStateDetails `yaml:"state"`
	Context  string           `yaml:"context"`
}

type workFlowDetails struct {
	Name      string        `yaml:"name"`
	StartTask string        `yaml:"startTask"`
	Task      []taskDetails `yaml:"task"`
}

type flow struct {
	Workflow workFlowDetails `yaml:"workflow"`
}

func readYamlFile() ([]byte, error) {
	yamlFile, err := ioutil.ReadFile("updateManager.yaml")
	return yamlFile, err
}

func (f *flow) decodeWorkFlow() error {

	yamlFile, _ := readYamlFile()

	err := yaml.Unmarshal(yamlFile, f)
	return err
}

func main() {

	var f flow

	err := f.decodeWorkFlow()
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	fmt.Println(f)
}
