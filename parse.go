// Package gtmpl add gettext support to Go standard html/template package.
package gtmpl

import (
	"html/template"
	"os"
	"path/filepath"
)

func getAllFilePathsInDirectory(dirpath string) ([]string, error) {
	var paths []string
	err := filepath.Walk(dirpath, func(path string, info os.FileInfo, err error) error {
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

	return paths, nil
}

// Recursively parse all files in directory, including sub-directories.
func ParseDirectoryTree(domain, localeDir, dir string) (t *template.Template, err error) {
	paths, err := getAllFilePathsInDirectory(dir)
	if err != nil {
		return
	}
	return ParseFiles(domain, localeDir, paths...)
}

// Recursively parse all files in directory, including sub-directories with
// *gettext* function.
//
// *gettext* function will translate input string according to installed
// translations and locale.
func ParseFiles(domain, localeDir string, filenames ...string) (*template.Template, error) {
	setupMessagesDomain(domain, localeDir)
	funcMap := template.FuncMap{
		"gettext": Translate,
	}
	return template.New(filenames[0]).Funcs(funcMap).ParseFiles(filenames...)
}
