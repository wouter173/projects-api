package misc

import (
	"strings"

	"github.com/wouter173/projects/structs"
	"gopkg.in/yaml.v2"
)

func GetMeta(content string) structs.Meta {

	str := strings.Split(string(content), "---")

	var meta structs.Meta
	yaml.Unmarshal([]byte(str[1]), &meta)

	return meta
}
