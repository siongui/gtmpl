// Some utilities for Go template
package gossg

import (
	"html/template"
	"os"
	"path/filepath"
)

// Recursively parse all files in directory, including sub-directories.
func ParseDirectory(dir string) (*template.Template, error) {
	var paths []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}
