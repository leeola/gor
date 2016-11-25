package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/leeola/gor/gobin"
)

func main() {
	var (
		source string
		args   []string
	)
	if len(os.Args) > 1 {
		source = os.Args[1]
		args = os.Args[2:]
	}

	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	output, err := makeOutputPath(source)
	if err != nil {
		panic(err)
	}

	b, err := gobin.New(gobin.Config{
		Cwd:    cwd,
		Stderr: os.Stderr,
	})

	exit, err := b.Build(gobin.BuildConfig{
		Source: source,
		Output: output,
	})
	if err != nil {
		panic(err)
	}

	if exit != 0 {
		os.Exit(0)
	}

	cmd := exec.Command("./"+output, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				os.Exit(status.ExitStatus())
			}
		}
		// If it's not an execerr or we can't get the status, panic err
		panic(err)
	}
}

func makeOutputPath(sourcePath string) (string, error) {
	// if it is a dir (go package), the produced binary will be the name
	// of the dir/package (with the gor extension)
	output := filepath.Base(sourcePath)

	// If it's a file, the produced binary will be the name of the file (with the gor
	// extension)
	if ext := filepath.Ext(output); ext != "" {
		output = strings.TrimSuffix(output, ext)
	}

	return output + ".gor", nil
}
