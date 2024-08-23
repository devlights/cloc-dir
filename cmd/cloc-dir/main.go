package main

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
)

type (
	Args struct {
		dir  string
		sep  string
		lang string
		cloc string
	}
)

var (
	args Args
)

func init() {
	flag.StringVar(&args.dir, "dir", ".", "directory")
	flag.StringVar(&args.sep, "sep", ",", "separator")
	flag.StringVar(&args.lang, "lang", "C#", "Language")
	flag.StringVar(&args.cloc, "cloc", "cloc", "cloc executable path")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if args.dir == "" {
		args.dir = "."
	}

	if args.sep == "" || len(args.sep) > 1 {
		args.sep = ","
	}

	if args.cloc == "" {
		args.cloc = "cloc"
	}

	if strings.HasPrefix(args.cloc, "./") {
		args.cloc = abs(args.cloc)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	entries, err := os.ReadDir(args.dir)
	if err != nil {
		return err
	}

	writer := csv.NewWriter(os.Stdout)
	writer.Comma = rune(args.sep[0])
	defer writer.Flush()

	columns := []string{"dir", "files", "language", "blank", "comment", "code"}
	if err := writer.Write(columns); err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		target := filepath.Join(args.dir, entry.Name())
		out, err := exec.Command(args.cloc, "--quiet", "--csv", "--hide-rate", fmt.Sprintf("--include-lang=%s", args.lang), target).Output()
		if err != nil {
			return err
		}

		buf := bytes.NewBuffer(out)
		scanner := bufio.NewScanner(buf)
		lines := make([]string, 0, 32)

		for scanner.Scan() {
			line := scanner.Text()
			lines = append(lines, line)
		}

		szLines := len(lines)
		for i, line := range lines {
			if i == 0 || i == szLines-1 {
				continue
			}

			parts := strings.Split(line, ",")
			parts = slices.Insert(parts, 0, entry.Name())
			if err := writer.Write(parts); err != nil {
				return err
			}
		}

		writer.Flush()
	}

	return nil
}

func abs(p string) string {
	v, _ := filepath.Abs(p)
	return v
}
