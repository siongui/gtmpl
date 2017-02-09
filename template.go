// Some utilities for Go template
package gossg

import (
	"html/template"
	"os"
	"path/filepath"
)

// Recursively get all file paths in directory, including sub-directories.
func GetAllFilePathsInDirectory(dirpath string) ([]string, error) {
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
func ParseDirectory(dirpath string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return nil, err
	}
	return template.ParseFiles(paths...)
}

// Recursively parse all files in directory, including sub-directories with
// *gettext* function.
//
// *gettext* function will translate input string according to installed
// translations and locale.
func ParseDirectoryWithGettextFunction(dirpath string) (*template.Template, error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return nil, err
	}

	funcMap := template.FuncMap{
		"gettext": Translate,
	}

	return template.New("").Funcs(funcMap).ParseFiles(paths...)
}
