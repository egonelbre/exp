package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/sync/errgroup"
)

var versions = []string{
	"go1.7.6",
	"go1.8.7",
	"go1.9.7",
	"go1.10.8",
	"go1.11.13",
	"go1.12.17",
	"go1.13.15",
	"go1.14.15",
	"go1.15.15",
	"go1.16.15",
	"go1.17.13",
	"go1.18.10",
	"go1.19.12",
	"go1.20.7",
	"go1.21.0",
}

func main() {
	var g errgroup.Group
	g.SetLimit(4)
	for _, version := range versions {
		version := version
		// if _, err := exec.LookPath(version); err == nil {
		// 	continue
		// }
		g.Go(func() error {
			fmt.Println("Installing", version)
			must0(run("go", "install", "golang.org/dl/"+version+"@latest"))
			must0(run(version, "download"))
			return nil
		})
	}
	g.Wait()

	for _, version := range versions {
		fmt.Println("# Benchmarking ", version)
		runtee(version+".log", version, "test", "-bench=.", "-count=6", ".")
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	return v
}

func must0(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(cmd string, args ...string) error {
	fmt.Println("# ", cmd, strings.Join(args, " "))
	x := exec.Command(cmd, args...)
	x.Stderr = os.Stderr
	x.Stdout = os.Stdout
	return x.Run()
}

func runtee(outfile string, cmd string, args ...string) error {
	fmt.Println("# ", cmd, strings.Join(args, " "))
	x := exec.Command(cmd, args...)
	f, err := os.Create(outfile)
	if err != nil {
		return err
	}
	defer f.Close()

	bf := bufio.NewWriter(f)
	defer bf.Flush()

	x.Stderr = os.Stderr
	x.Stdout = io.MultiWriter(os.Stdout, bf)
	return x.Run()
}

/*
#!/bin/bash

declare -a versions=(
)

echo "# Installing Go versions"

for ver in "${versions[@]}"
do
	go install "golang.org/dl/${ver}@latest"
	${ver} download
done

echo "# Benchmarking"

for ver in "${versions[@]}"
do
	$
{ver} test -bench . -count 6 | tee ${ver}.log
done
*/
