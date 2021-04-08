package handlers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/wouter173/projects/misc"
	"github.com/wouter173/projects/structs"
)

func IDMetaHandler(c *fiber.Ctx) error {
	c.Set("content-type", "application/json")
	c.Set("Access-Control-Allow-Origin", "*")

	meta, err := misc.GetMeta(c.Params("id"))

	if err != nil {
		res, _ := json.Marshal(structs.Error{Message: "Project not found."})
		return c.Send(res)
	}

	res, _ := json.Marshal(meta)
	return c.Send(res)
}
