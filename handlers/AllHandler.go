package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/wouter173/projects/misc"
	"github.com/wouter173/projects/structs"
)

//AllHandler get all projects in defined path
func AllHandler(c *fiber.Ctx) error {
	metas, err := misc.GetAllMeta()

	c.Set("content-type", "application/json")
	if err != nil {
		res, _ := json.Marshal(structs.Error{Message: err.Error()})
		return c.Send(res)
	}

	res, _ := json.Marshal(metas)

	return c.Send(res)
}
