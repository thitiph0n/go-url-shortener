package repositories

import "time"

//Type: generated/custom

type Link struct {
	LinkId    string    `json:"linkId" firestore:"linkId"`
	Url       string    `json:"url" firestore:"url"`
	Clicked   int       `json:"clicked" firestore:"clicked"`
	Type      string    `json:"type" firestore:"type"`
	CreatedAt time.Time `json:"createdAt" firestore:"createdAt"`
}

type LinkRepository interface {
	GetAll() ([]Link, error)
	GetById(string) (*Link, error)
	GetByUrl(string) (*Link, error)
	Create(Link) error
	Update(Link) error
}
