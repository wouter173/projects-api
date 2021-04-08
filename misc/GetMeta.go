package misc

import (
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/wouter173/projects/structs"
	"gopkg.in/yaml.v2"
)

func GetMeta(name string) (structs.Meta, error) {

	content, err := ioutil.ReadFile(os.Getenv("PROJECTS_PATH") + "/" + name + ".md")

	if err != nil {
		return structs.Meta{}, errors.New("Project not found.")
	}

	str := strings.Split(string(content), "---")

	if len(str) == 1 {
		return structs.Meta{Name: name}, nil
	}

	var meta structs.Meta
	yaml.Unmarshal([]byte(str[1]), &meta)

	return meta, nil
}

func GetAllMeta() ([]structs.Meta, error) {

	var files []structs.Meta
	dir, err := ioutil.ReadDir(os.Getenv("PROJECTS_PATH"))

	if err != nil {
		return nil, errors.New("Path not found.")
	}

	for _, file := range dir {
		project, _ := GetMeta(strings.TrimRight(file.Name(), ".md"))
		files = append(files, project)
	}

	return files, nil
}
