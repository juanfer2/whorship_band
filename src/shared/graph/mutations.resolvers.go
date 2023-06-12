package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

import (
	"context"

	"github.com/juanfer2/whorship_band/src/shared/graph/model"
	songs_application "github.com/juanfer2/whorship_band/src/songs/application"
	songs_domain "github.com/juanfer2/whorship_band/src/songs/domain"
	songs_repositories "github.com/juanfer2/whorship_band/src/songs/infrastructure/repositories"
)

// CreateSong is the resolver for the createSong field.
func (r *mutationResolver) CreateSong(ctx context.Context, input model.SongInput) (*songs_domain.Song, error) {
	songRepository := songs_repositories.NewSongPGRepository()
	songUseCase := songs_application.NewSongUseCase(songRepository)
	song, err := songUseCase.Create(songs_domain.Song{
		Name: input.Name,
		Url:  *input.URL,
		Tone: *input.Tone,
	})

	return &song, err
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
