package models

import (
	"io/ioutil"
	"os"
	"os/exec"
)

// Code is the script code
type Code string

func (c Code) dirPath(name string) string {
	return "codes/" + name
}

func (c Code) filePath(name string) string {
	return c.dirPath(name) + "/main.go"
}

func (c Code) pluginPath(name string) string {
	return c.dirPath(name) + "/main.so"
}

// Configure saves code to a file and builds plugin
func (c Code) Configure(name string) error {
	var err error
	err = os.MkdirAll(c.dirPath(name), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(c.filePath(name), []byte(c), 0644)
	if err != nil {
		return err
	}

	return exec.Command(
		"go", "build",
		"-buildmode", "plugin",
		"-o", c.pluginPath(name),
		c.filePath(name),
	).Run()
}
