package engine

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Engine struct {
}

func (e *Engine) Gen(input string) string {
	return "package main\n\nfunc main() {\n" + input + "\n}"
}

func (e *Engine) Save(code string) (string, string) {
	dir, err := ioutil.TempDir("", "go-play-engine-temp")
	if err != nil {
		log.Fatal(err)
	}
	tmpfn := filepath.Join(dir, "main.go")
	if err := ioutil.WriteFile(tmpfn, []byte(code), 0666); err != nil {
		log.Fatal(err)
	}
	return dir, tmpfn
}

func (e *Engine) CleanUp(dir string) {
	os.RemoveAll(dir)
}

func (e *Engine) Run(code string) {
	path, _ := e.Save(e.Gen(code))
	//defer e.CleanUp(path)

	print("Executing... \n")

	os.Chdir(path)

	cmdName := "go"
	cmdArgs := []string{"run", "main.go"}
	cmd := exec.Command(cmdName, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}
