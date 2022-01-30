package repositories

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

type linkRepositoryFirestore struct {
	client *firestore.Client
	ctx    context.Context
}

func NewLinkRepositoryFirestore(ctx context.Context, client *firestore.Client) LinkRepository {
	return &linkRepositoryFirestore{client, ctx}
}

func (r *linkRepositoryFirestore) GetAll() ([]Link, error) {
	links := []Link{}

	iter := r.client.Collection("links").Documents(r.ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}

		var link Link
		doc.DataTo(&link)
		links = append(links, link)
	}

	return links, nil
}

func (r *linkRepositoryFirestore) GetById(id string) (*Link, error) {

	doc, err := r.client.Collection("links").Doc(id).Get(r.ctx)
	if err != nil {

		if err == iterator.Done {
			return nil, nil
		}

		return nil, err
	}

	var link Link
	doc.DataTo(&link)

	return &link, nil
}

func (r *linkRepositoryFirestore) GetByUrl(url string) (*Link, error) {

	doc, err := r.client.Collection("links").Where("url", "==", url).Limit(1).Documents(r.ctx).Next()
	if err != nil {
		if err == iterator.Done {
			return nil, nil
		}

		return nil, err
	}

	var link Link
	doc.DataTo(&link)

	return nil, nil
}

func (r *linkRepositoryFirestore) Create(link Link) error {

	_, err := r.client.Collection("links").Doc(link.LinkId).Set(r.ctx, link)
	if err != nil {
		return err
	}

	return nil
}

func (r *linkRepositoryFirestore) Update(link Link) error {

	_, err := r.client.Collection("links").Doc(link.LinkId).Set(r.ctx, link)
	if err != nil {
		return err
	}

	return nil
}
