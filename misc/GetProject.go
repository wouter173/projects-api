package misc

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/wouter173/projects/structs"
)

func GetProject(name string) (structs.Meta, error) {

	content, err := ioutil.ReadFile(os.Getenv("PROJECTS_PATH") + "/" + name + ".md")

	if err != nil {
		return structs.Meta{}, errors.New("Project not found.")
	}

	return GetMeta(string(content)), nil
}

func GetAllProjects() ([]structs.Meta, error) {

	var files []structs.Meta
	dir, err := ioutil.ReadDir(os.Getenv("PROJECTS_PATH"))

	if err != nil {
		return nil, errors.New("Path not found.")
	}

	for _, file := range dir {
		project, _ := GetProject(strings.TrimRight(file.Name(), ".md"))
		files = append(files, project)
	}

	return files, nil
}
