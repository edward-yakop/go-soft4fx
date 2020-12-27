package path

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func ResultHtmlFiles() []string {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("Missing folder/file argument")
	}

	keys := make(map[string]bool)
	var results []string
	for _, arg := range args {
		if fi, err := os.Stat(arg); err != nil {
			log.Println("[", arg, "] is not valid. Error: ", err)
			continue
		} else if fi.IsDir() {
			_ = filepath.Walk(arg, func(path string, info os.FileInfo, err error) error {
				if isHtmlFileName(path) {
					if _, value := keys[path]; !value {
						keys[path] = true
						results = append(results, path)
					}
				}
				return nil
			})
		} else if isHtmlFileName(arg) {
			if _, value := keys[arg]; !value {
				keys[arg] = true
				results = append(results, arg)
			}
		}
	}

	return results
}

func isHtmlFileName(path string) bool {
	return strings.HasSuffix(path, ".htm") || strings.HasSuffix(path, ".html")
}
