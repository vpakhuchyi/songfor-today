package datastore

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"

	"github.com/vpakhuchyi/songfor-today/models"
)

const kindTrack = "Track"

type Adapter struct {
	ds *datastore.Client
}

func New(ds *datastore.Client) Adapter {
	return Adapter{ds: ds}
}

func (a *Adapter) PutTracks(ctx context.Context, tracks []models.Track) error {
	keys := make([]*datastore.Key, len(tracks))
	for i := 0; i < len(tracks); i++ {
		keys[i] = datastore.IDKey(kindTrack, int64(tracks[i].ID), nil)
	}

	if _, err := a.ds.PutMulti(ctx, keys, tracks); err != nil {
		return fmt.Errorf("failed to save tracks: %w", err)
	}

	return nil
}

func (a *Adapter) GetAllTracks(ctx context.Context) ([]models.Track, error) {
	q := datastore.NewQuery(kindTrack)

	var tracks []models.Track
	if _, err := a.ds.GetAll(ctx, q, tracks); err != nil {
		return nil, fmt.Errorf("failed to get all tracks: %w", err)
	}

	return tracks, nil
}

func (a *Adapter) GetTrack(ctx context.Context, id int) (models.Track, error) {
	key := datastore.IDKey(kindTrack, int64(id), nil)

	var track models.Track
	if err := a.ds.Get(ctx, key, &track); err != nil {
		return models.Track{}, fmt.Errorf("failed to get track: %w", err)
	}

	return track, nil
}
