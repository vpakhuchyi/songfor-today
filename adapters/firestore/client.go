package firestore

import (
	"context"
	"fmt"
	"strconv"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/vpakhuchyi/songfor-today/models"
)

const collectionTracks = "tracks"

type Adapter struct {
	fs *firestore.Client
}

func New(fs *firestore.Client) Adapter {
	return Adapter{fs: fs}
}

func (a *Adapter) PutTracks(ctx context.Context, tracks []models.Track) error {
	for i := 0; i < len(tracks); i++ {
		_, err := a.fs.Collection(collectionTracks).Doc(strconv.Itoa(tracks[i].ID)).Set(ctx, tracks[i])
		if err != nil {
			return fmt.Errorf("failed to put track id=%d: %w", tracks[i].ID, err)
		}
	}

	return nil
}

func (a *Adapter) GetAllTracks(ctx context.Context) ([]models.Track, error) {
	var tracks []models.Track

	iter := a.fs.Collection(collectionTracks).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to get tracks iterator: %w", err)
		}

		var track models.Track
		err = doc.DataTo(&track)
		if err != nil {
			return nil, fmt.Errorf("failed to decode a track: %w", err)
		}

		tracks = append(tracks, track)
	}

	return tracks, nil
}
