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

	files := getFiles()

	jobs := make(chan string, len(files))
	results := make(chan bool, len(files))

	for w := 1; w <= 20; w++ {
		go worker(w, jobs, results)
	}

	for _, file := range files {
		jobs <- file
	}

	close(jobs)

	for a := 1; a <= len(files); a++ {
		<-results
	}
}

func getFiles() []string {
	var files []string

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		path = strings.TrimSpace(path)
		if isLintable(path) {
			files = append(files, path)
		}
		return err
	})

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
