package repositories

//Type: generated/custom

type Link struct {
	LinkId  string `json:"linkId"`
	Url     string `json:"url"`
	Clicked int    `json:"clicked"`
	Type    string `json:"type"`
}

type LinkRepository interface {
	GetAll() ([]Link, error)
	GetById(string) (*Link, error)
	GetByUrl(string) (*Link, error)
	Create(Link) error
	Update(Link) error
}
