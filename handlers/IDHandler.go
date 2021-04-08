package handlers

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/wouter173/projects/structs"
)

//IDHandler get 1 project from defined path and print it
func IDHandler(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")

	file, err := ioutil.ReadFile(os.Getenv("PROJECTS_PATH") + "/" + c.Params("id") + ".md")

	if err != nil {
		c.Set("Content-type", "application/json")
		res, _ := json.Marshal(structs.Error{Message: "Project not found."})
		return c.Send(res)
	}

	str := strings.Split(string(file), "---")

	if len(str) == 1 {
		return c.SendString(str[0])
	}

	return c.SendString(str[2])
}
