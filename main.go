package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	yaml "gopkg.in/yaml.v2"
)

type Cmd struct {
	Name    string   `yaml:"name"`
	Command string   `yaml:"command"`
	Async   bool     `yaml:"async,omitempty"`
	CanFail bool     `yaml:"can_fail,omitempty"`
	Deps    []string `yaml:"deps,omitempty"`
}

type Dag []Cmd

func execute(command string, name string) {
	cmd := exec.Command("bash", "-c", command)
	fmt.Printf("[%s]\n%s\n", name, cmd)
	out, err := cmd.Output()
	handleErr(err)
	fmt.Println(string(out))
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func parseInput() Dag {
	f, err := os.Open("./examples/input.yml")
	handleErr(err)

	bytes, err := ioutil.ReadAll(f)
	handleErr(err)

	inp := make(Dag, 0)
	err = yaml.Unmarshal(bytes, &inp)
	handleErr(err)
	return inp
}

func main() {
	dag := parseInput()
	fmt.Println(dag)

	for _, cmd := range dag {
		execute(cmd.Command, cmd.Name)
	}
}
