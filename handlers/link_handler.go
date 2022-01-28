package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/thitiph0n/go-url-shortener/services"
)

type linkHandler struct {
	linkService services.LinkService
}

func NewLinkHandler(linkService services.LinkService) LinkHandler {
	return linkHandler{linkService}
}

func (h linkHandler) CreateLink(c *fiber.Ctx) error {
	return nil
}

func (h linkHandler) GetLinks(c *fiber.Ctx) error {
	return nil
}

func (h linkHandler) GetLinkById(c *fiber.Ctx) error {
	return nil
}

func (h linkHandler) ResloveLink(c *fiber.Ctx) error {
	return nil
}
