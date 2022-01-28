package repositories

import "cloud.google.com/go/firestore"

type linkRepositoryFirestore struct {
	client *firestore.Client
}

func NewLinkRepositoryFirestore(client *firestore.Client) LinkRepository {
	return &linkRepositoryFirestore{client}
}

func (r *linkRepositoryFirestore) GetAll() ([]Link, error) {
	return nil, nil
}

func (r *linkRepositoryFirestore) GetById(id string) (*Link, error) {
	return nil, nil
}

func (r *linkRepositoryFirestore) GetByUrl(url string) (*Link, error) {
	return nil, nil
}

func (r *linkRepositoryFirestore) Create(link Link) error {
	return nil
}

func (r *linkRepositoryFirestore) Update(link Link) error {
	return nil
}
