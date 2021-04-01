package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
)

//Error struct
type Error struct {
	Message string
}

//AllHandler get all projects in defined path
func AllHandler(c *fiber.Ctx) error {
	c.Set("content-type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")

	var files []string
	dir, err := ioutil.ReadDir(os.Getenv("PROJECTS_PATH"))

	if err != nil {
		res, _ := json.Marshal(Error{Message: "Path not found."})
		return c.Send(res)
	}

	for _, file := range dir {
		files = append(files, strings.TrimRight(file.Name(), ".md"))
	}

	res, _ := json.Marshal(files)

	return c.Send(res)
}

//IDHandler get 1 project from defined path and print it
func IDHandler(c *fiber.Ctx) error {
	c.Set("Access-Control-Allow-Origin", "*")

	file, err := ioutil.ReadFile(os.Getenv("PROJECTS_PATH") + "/" + c.Params("id") + ".md")

	if err != nil {
		c.Set("Content-type", "application/json")
		res, _ := json.Marshal(Error{Message: "Project not found."})
		return c.Send(res)
	}

	return c.SendString(string(file))
}
