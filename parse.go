// Package gtmpl adds gettext support to Go standard html/template package.
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

// ParseDirectoryTree will find all files in directory tree (including
// sub-directories) and pass filenames to ParseFiles.
func ParseDirectoryTree(domain, localeDir, dir string) (t *template.Template, err error) {
	paths, err := getAllFilePathsInDirectory(dir)
	if err != nil {
		return
	}
	return ParseFiles(domain, localeDir, paths...)
}

// ParseFiles is the same as ParseFiles except gettext function is installed via
// Funcs method provided by Go standard html/template package. the installed
// gettext function will use the data given by domain and localeDir arguments to
// translate the strings in the template.
func ParseFiles(domain, localeDir string, filenames ...string) (*template.Template, error) {
	setupMessagesDomain(domain, localeDir)
	funcMap := template.FuncMap{
		"gettext": Translate,
	}
	return template.New(filenames[0]).Funcs(funcMap).ParseFiles(filenames...)
}
