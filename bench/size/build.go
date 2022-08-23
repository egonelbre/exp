package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	entries, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, e := range entries {
		if !e.IsDir() {
			continue
		}

		outfile := e.Name() + ".exe"
		cmd := exec.Command("go1.19", "build", "-o", outfile, "./"+e.Name())
		cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Fprintln(os.Stderr, string(out))
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		stat, err := os.Stat(outfile)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		fmt.Println(outfile, stat.Size())
	}
}
