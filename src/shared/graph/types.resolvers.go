package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"

	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
)

// ID is the resolver for the id field.
func (r *songResolver) ID(ctx context.Context, obj *songs_domain.Song) (*string, error) {
	id := obj.ID.String()

	return &id, nil
}

// Song returns SongResolver implementation.
func (r *Resolver) Song() SongResolver { return &songResolver{r} }

type songResolver struct{ *Resolver }