// Go template manager
package gotm

import (
	"html/template"
	"io"
	"os"
	"path/filepath"
)

type TemplateManager struct {
	name    string
	StdTmpl *template.Template
}

func NewTemplateManager(name string) *TemplateManager {
	return &TemplateManager{
		name: name,
	}
}

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
func (tm *TemplateManager) ParseDirectory(dirpath string) (err error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return
	}
	tm.StdTmpl, err = template.ParseFiles(paths...)
	return
}

// Recursively parse all files in directory, including sub-directories with
// *gettext* function.
//
// *gettext* function will translate input string according to installed
// translations and locale.
func (tm *TemplateManager) ParseDirectoryWithGettextFunction(dirpath string) (err error) {
	paths, err := GetAllFilePathsInDirectory(dirpath)
	if err != nil {
		return
	}

	funcMap := template.FuncMap{
		"gettext": Translate,
	}

	tm.StdTmpl, err = template.New(tm.name).Funcs(funcMap).ParseFiles(paths...)
	return
}

// Same paramaters as *template.ExecuteTemplate
func (tm *TemplateManager) ExecuteTemplate(wr io.Writer, name string, data interface{}) error {
	return tm.StdTmpl.ExecuteTemplate(wr, name, data)
}
