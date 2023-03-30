package handlers

import (
	"context"
	"fmt"
	"html/template"

	"github.com/gofiber/fiber/v2"
	gowiki "github.com/trietmn/go-wiki"
	"golang.org/x/exp/slog"

	"github.com/vpakhuchyi/songfor-today/config"
	"github.com/vpakhuchyi/songfor-today/models"
)

type TracksProvider interface {
	RandomSong(ctx context.Context) (models.Track, error)
}

type Tracks struct {
	tp  TracksProvider
	cfg config.Config
}

func NewTracks(tp TracksProvider, config config.Config) Tracks {
	return Tracks{
		tp:  tp,
		cfg: config,
	}
}

func (t *Tracks) RandomSong(ctx *fiber.Ctx) error {
	track, err := t.tp.RandomSong(ctx.Context())
	if err != nil {
		return fmt.Errorf("failed to get a random song: %w", err)
	}

	slog.Info("Random song", "track", track)
	//
	page, err := gowiki.GetPage(track.Artist.Name, -1, false, true)
	if err != nil {
		return fmt.Errorf("failed to get a data from wiki: %w", err)
	}

	// Get the content of the page
	content, err := page.GetSummary()
	if err != nil {
		return fmt.Errorf("failed to extract wiki content: %w", err)
	}

	data := struct {
		AlbumImage    string
		Artist        string
		Biography     string
		BiographyLink string
		Track         string
		TrackLink     string
		TrackPreview  string
	}{
		AlbumImage:    track.Album.CoverBig,
		Artist:        track.Artist.Name,
		Biography:     content,
		BiographyLink: page.URL,
		Track:         track.Title,
		TrackLink:     track.Link,
		TrackPreview:  track.Preview,
	}

	// Parse the HTML template.
	tmpl, err := template.ParseFiles("./static/random-song.tmpl")
	if err != nil {
		return fmt.Errorf("failed to parse a html template: %w", err)
	}

	ctx.Response().Header.Set("Content-Type", "text/html")

	// Execute the template with the data.
	err = tmpl.Execute(ctx.Response().BodyWriter(), data)
	if err != nil {
		return fmt.Errorf("failed to execute a html template: %w", err)
	}

	//
	//err = ctx.JSON(track)
	//if err != nil {
	//	return fmt.Errorf("failed to marshal a song: %w", err)
	//}

	return nil
}
