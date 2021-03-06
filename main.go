// command phpl is a wrapper for php's native linter which adds support for
// linting many files concurrently.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	err := exec.Command("which", "php").Run()
	if err != nil {
		log.Fatal("command not found: php")
	}

	files := getFiles(".")
	length := len(files)

	jobs := make(chan string, length)
	results := make(chan bool, length)

	for w := 1; w <= 10; w++ {
		go worker(w, jobs, results)
	}

	for _, file := range files {
		jobs <- file
	}

	close(jobs)

	for a := 1; a <= length; a++ {
		<-results
	}
}

func getFiles(root string) []string {
	var files []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		path = strings.TrimSpace(path)
		if isLintable(path) {
			files = append(files, path)
		}
		return err
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func isLintable(path string) bool {
	if path == "" || strings.HasPrefix(path, "vendor") || strings.HasPrefix(path, "node_modules") {
		return false
	}
	return strings.HasSuffix(path, ".php")
}

func worker(id int, jobs <-chan string, results chan<- bool) {
	for j := range jobs {
		lint(j)
		results <- true
	}
}

func lint(file string) {
	stdout, err := exec.Command("php", "-l", file).Output()
	if err != nil {
		message := string(stdout)
		message = strings.TrimSpace(message)
		fmt.Println(message)
	}
}
