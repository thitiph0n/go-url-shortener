package handlers

import "github.com/gofiber/fiber/v2"

type LinkHandler interface {
	CreateLink(c *fiber.Ctx) error
	GetLinks(c *fiber.Ctx) error
	GetLinkById(c *fiber.Ctx) error
	ResloveLink(c *fiber.Ctx) error
}
