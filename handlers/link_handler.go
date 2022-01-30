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

	linkRequest := services.NewLinkRequest{}

	if err := c.BodyParser(&linkRequest); err != nil {
		return err
	}

	linkResponse, err := h.linkService.CreateLink(linkRequest)
	if err != nil {
		return err
	}

	return c.Status(201).JSON(linkResponse)
}

func (h linkHandler) GetLinks(c *fiber.Ctx) error {

	linkResponses, err := h.linkService.GetLinks()
	if err != nil {
		return err
	}

	return c.JSON(linkResponses)
}

func (h linkHandler) GetLinkById(c *fiber.Ctx) error {
	linkId := c.Params("linkId")
	linkResponse, err := h.linkService.GetLinkById(linkId)
	if err != nil {
		return err
	}

	return c.JSON(linkResponse)
}

func (h linkHandler) ResloveLink(c *fiber.Ctx) error {
	linkId := c.Params("linkId")
	linkResponse, err := h.linkService.ResloveLink(linkId)
	if err != nil {
		return err
	}

	return c.JSON(linkResponse)
}
