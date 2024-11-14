package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
)

var reset = "\033[0m"
var red = "\033[31m"
var green = "\033[32m"
var yellow = "\033[33m"

func main() {
	args := os.Args[1:]
	if 2 != len(args) {
		panic(fmt.Sprintf("%sgo-create should be called with two arguments, %d found%s", red, len(args), reset))
	}

	name := args[0]
	version := args[1]

	if !validateVersion(version) {
		panic(fmt.Sprintf("%sInvalid Go version: %s%s", red, version, reset))
	}

	if !validateName(name) {
		panic(fmt.Sprintf("%sInvalid application name: %s%s", red, name, reset))
	}

	err := os.CopyFS("../"+name, os.DirFS("src"))
	if nil != err {
		panic(err)
	}

	fs.WalkDir(os.DirFS("../"+name), ".", func(path string, d fs.DirEntry, err error) error {
		if nil != err {
			panic(err)
		}

		if d.IsDir() {
			return nil
		}

		b, err := os.ReadFile("../" + name + "/" + path)
		if nil != err {
			panic(err)
		}

		if "go.mod" == d.Name() {
			b = append(b, []byte("\ngo "+version)...)
		}

		err = os.WriteFile("../"+name+"/"+path, []byte(strings.Replace(string(b), "!!placeholder!!", "newApp", -1)), 0666)
		if nil != err {
			panic(err)
		}

		return nil
	})

	fmt.Println("")

	fmt.Printf("%sApplication %s created using Go version %s%s\n", green, name, version, reset)

	fmt.Println("")

	fmt.Println(yellow + "To run locally:" + reset)
	fmt.Println(" cd newApp")
	fmt.Println(" go mod tidy")
	fmt.Println(" go build && ./newApp")

	fmt.Println("")

	fmt.Println(yellow + "To run tests:" + reset)
	fmt.Println(" go test ./...")

	fmt.Println("")
}
