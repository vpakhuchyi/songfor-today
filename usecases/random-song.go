package usecases

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/vpakhuchyi/songfor-today/adapters/firestore"
	"github.com/vpakhuchyi/songfor-today/models"
)

type Client struct {
	ds *firestore.Adapter
}

func New(ds *firestore.Adapter) Client {
	return Client{ds: ds}
}

func (c Client) RandomSong(ctx context.Context) (models.Track, error) {
	tracks, err := c.ds.GetAllTracks(ctx)
	if err != nil {
		return models.Track{}, fmt.Errorf("failed to fetch tracks: %w", err)
	}

	return tracks[rand.Intn(len(tracks))], nil
}
