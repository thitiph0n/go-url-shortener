package services

import "github.com/thitiph0n/go-url-shortener/repositories"

type linkService struct {
	linkRepo repositories.LinkRepository
}

func NewLinkService(linkRepo repositories.LinkRepository) LinkService {
	return linkService{linkRepo: linkRepo}
}

func (s linkService) CreateLink(NewLinkRequest) (*LinkResponse, error) {
	return nil, nil
}

func (s linkService) GetLinkById(string) (*LinkResponse, error) {
	return nil, nil
}

func (s linkService) ResloveLink(string) (*LinkResponse, error) {
	return nil, nil
}

func (s linkService) GetLinks() ([]LinkResponse, error) {
	return nil, nil
}
