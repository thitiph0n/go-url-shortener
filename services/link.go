package services

type NewLinkRequest struct {
	Url          string `json:"url"`
	CustomLinkId string `json:"customLinkId,omitempty"`
}

type LinkResponse struct {
	LinkId  string `json:"linkId"`
	Url     string `json:"url"`
	Clicked int    `json:"clicked"`
}

type LinkService interface {
	CreateLink(NewLinkRequest) (*LinkResponse, error)
	GetLinkById(string) (*LinkResponse, error)
	ResloveLink(string) (*LinkResponse, error)
	GetLinks() ([]LinkResponse, error)
}
