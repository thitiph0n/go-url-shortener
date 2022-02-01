package services

import (
	"log"
	"time"

	"github.com/thitiph0n/go-url-shortener/errs"
	"github.com/thitiph0n/go-url-shortener/helpers"
	"github.com/thitiph0n/go-url-shortener/repositories"
)

type linkService struct {
	linkRepo repositories.LinkRepository
}

func NewLinkService(linkRepo repositories.LinkRepository) LinkService {
	return linkService{linkRepo: linkRepo}
}

func (s linkService) CreateLink(linkRequest NewLinkRequest) (*LinkResponse, error) {

	// check is valid url
	if !helpers.CheckDomainError(linkRequest.Url) {
		return nil, errs.NewBadRequestError("invalid url")
	}

	// check is link exist
	exist, err := s.linkRepo.GetByUrl(linkRequest.Url)
	if err != nil {
		return nil, err
	}

	if exist != nil {
		return &LinkResponse{
			Url:     exist.Url,
			LinkId:  exist.LinkId,
			Clicked: exist.Clicked,
		}, nil
	}

	link := repositories.Link{}

	link.Url = linkRequest.Url

	// check type of link
	if linkRequest.CustomLinkId == "" {
		link.Type = Generated.String()
		link.Clicked = 0
		link.CreatedAt = time.Now().UTC()
		link.LinkId = helpers.GenerateLinkId(6)

		if s.linkRepo.Create(link) != nil {
			return nil, errs.NewUnexpectedError()
		}

		return &LinkResponse{
			Url:     link.Url,
			LinkId:  link.LinkId,
			Clicked: link.Clicked,
		}, nil
	}

	link.Type = Custom.String()

	// check custom link is valid
	if len(linkRequest.CustomLinkId) > 32 {
		return nil, errs.NewBadRequestError("invalid custom linkId")
	}

	// check if custom link is exist
	if exist, err := s.linkRepo.GetById(linkRequest.CustomLinkId); err != nil && exist != nil {
		return nil, errs.NewBadRequestError("custom linkId already exist")
	}

	link.LinkId = linkRequest.CustomLinkId
	link.Clicked = 0
	link.CreatedAt = time.Now().UTC()

	if s.linkRepo.Create(link) != nil {
		return nil, errs.NewUnexpectedError()
	}

	return &LinkResponse{
		Url:     link.Url,
		LinkId:  link.LinkId,
		Clicked: link.Clicked,
	}, nil
}

func (s linkService) GetLinkById(linkId string) (*LinkResponse, error) {
	link, err := s.linkRepo.GetById(linkId)
	if err != nil {
		log.Printf("[Service] GetLinkById: %v", err)
		return nil, errs.NewUnexpectedError()
	}

	if link == nil {
		return nil, errs.NewNotFoundError("link not found")
	}

	linkResponse := &LinkResponse{
		Url:     link.Url,
		LinkId:  link.LinkId,
		Clicked: link.Clicked,
	}

	return linkResponse, nil
}

func (s linkService) ResloveLink(linkId string) (*LinkResponse, error) {
	link, err := s.linkRepo.GetById(linkId)
	if err != nil {
		log.Printf("[Service] ResloveLink: %v", err)
		return nil, errs.NewUnexpectedError()
	}

	if link == nil {
		return nil, errs.NewNotFoundError("link not found")
	}

	link.Clicked += 1

	if s.linkRepo.Update(*link) != nil {
		return nil, errs.NewUnexpectedError()
	}

	linkResponse := &LinkResponse{
		Url:     link.Url,
		LinkId:  link.LinkId,
		Clicked: link.Clicked,
	}

	return linkResponse, nil
}

func (s linkService) GetLinks() ([]LinkResponse, error) {
	links, err := s.linkRepo.GetAll()
	if err != nil {
		log.Printf("[Service] GetLinks: %v", err)
		return nil, errs.NewUnexpectedError()
	}

	linkResponses := []LinkResponse{}

	for _, link := range links {
		linkResponse := LinkResponse{
			Url:     link.Url,
			LinkId:  link.LinkId,
			Clicked: link.Clicked,
		}
		linkResponses = append(linkResponses, linkResponse)
	}

	return linkResponses, nil
}
