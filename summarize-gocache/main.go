// gocachesummary parses build files in `go env GOCACHE` folder and summarizes
// how much space is used for a package builds.
package main

import (
	"bytes"
	"cmp"
	"flag"
	"fmt"
	"go/token"
	"go/types"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"text/tabwriter"

	"golang.org/x/tools/go/gcexportdata"
)

const MiB = 1 << 20

func main() {
	quiet := flag.Bool("quiet", false, "don't print parsing errors")
	minsize := flag.Int64("minsize", 512_000, "minimum file size that the tool cares about")
	flag.Parse()

	root := flag.Arg(0)
	if root == "" {
		fmt.Fprintln(os.Stderr, "please provide directory")
		os.Exit(1)
	}

	type Stat struct {
		Size  int64
		Count int64
	}

	statByPkgPath := map[string]Stat{}

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}
		if d.IsDir() {
			return nil
		}

		stat, err := d.Info()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return nil
		}

		if stat.Size() < *minsize {
			return nil
		}

		pkgpath, err := tryExtractPackagePath(path)
		if err != nil {
			if !*quiet {
				fmt.Fprintln(os.Stderr, err)
			}
			return nil
		}

		s := statByPkgPath[pkgpath]
		s.Size += stat.Size()
		s.Count++
		statByPkgPath[pkgpath] = s

		return nil
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	type Entry struct {
		PkgPath string
		Stat
	}

	var entries []Entry
	for pkgpath, stat := range statByPkgPath {
		entries = append(entries, Entry{
			PkgPath: pkgpath,
			Stat:    stat,
		})
	}

	slices.SortFunc(entries, func(a, b Entry) int {
		return cmp.Compare(a.Size, b.Size)
	})

	w := tabwriter.NewWriter(os.Stdout, 4, 4, 4, ' ', 0)
	defer w.Flush()

	fmt.Fprintf(w, "package\tsum (MiB)\tcount\tavg. size (MiB)\n")
	for _, entry := range entries {
		fmt.Fprintf(w, "%s\t%.2f\t%v\t%.2f\n", entry.PkgPath, float64(entry.Size)/MiB, entry.Count, (float64(entry.Size)/float64(entry.Count))/MiB)
	}
}

func must[T any](v T, err error) T {
	if err != nil {
		panic(err)
	}
	return v
}

func tryExtractPackagePath(path string) (pkgpath string, err error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("%s: unable to open package: %w", path, err)
	}

	r, err := gcexportdata.NewReader(bytes.NewReader(data))
	if err != nil {
		return "", fmt.Errorf("%s: not an export data file: %w", path, err)
	}

	fset := token.NewFileSet()
	imports := map[string]*types.Package{}

	pkg, err := gcexportdata.Read(r, fset, imports, "")
	if err != nil {
		return "", fmt.Errorf("%s: failed to read export data: %w", path, err)
	}

	return pkg.Path(), nil
}
