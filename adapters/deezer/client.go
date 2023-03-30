package deezer

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
	"golang.org/x/exp/slog"

	"github.com/vpakhuchyi/songfor-today/models"
)

type Client struct {
	httpclient *resty.Client
}

func New(httpclient *resty.Client) Client {
	return Client{httpclient: httpclient}
}

func (c *Client) Me(ctx context.Context, token string) error {
	resp, err := c.httpclient.SetDebug(true).
		R().
		SetContext(ctx).
		SetAuthToken(token).
		Get("https://api.deezer.com/user/me")
	if err != nil {
		return fmt.Errorf("failed to get user's deezer profile: %w", err)
	}

	slog.DebugCtx(ctx, "profile", resp.String())

	return nil
}

func (c *Client) GetPlaylists(ctx context.Context, token string) error {
	resp, err := c.httpclient.SetDebug(true).
		R().
		SetContext(ctx).
		SetAuthToken(token).
		Get("https://api.deezer.com/user/me/playlists")
	if err != nil {
		return fmt.Errorf("failed to get user's deezer playlists: %w", err)
	}

	slog.DebugCtx(ctx, "playlists", resp.String())

	return nil
}

func (c *Client) GetPlaylistTracks(ctx context.Context, token, playlistID string) ([]models.Track, error) {
	var data = struct {
		Data []models.Track `json:"data"`
	}{}

	resp, err := c.httpclient.SetDebug(true).
		R().
		SetContext(ctx).
		SetAuthToken(token).
		SetResult(&data).
		Get("https://api.deezer.com/playlist/" + playlistID + "/tracks")
	if err != nil {
		return nil, fmt.Errorf("failed to get a deezer playlist tracks: %w", err)
	}

	if !resp.IsSuccess() {
		return nil, fmt.Errorf("failed to get a playlist tracks: %s", resp.String())
	}

	slog.DebugCtx(ctx, "tracks", data)

	return data.Data, nil
}

func (c *Client) GetTrack(ctx context.Context, token, trackID string) (models.Track, error) {
	var track models.Track
	resp, err := c.httpclient.SetDebug(true).
		R().
		SetContext(ctx).
		SetAuthToken(token).
		SetResult(&track).
		Get("https://api.deezer.com/track/" + trackID)
	if err != nil {
		return models.Track{}, fmt.Errorf("failed to get a track: %w", err)
	}

	if !resp.IsSuccess() {
		return models.Track{}, fmt.Errorf("failed to get a track: %s", resp.String())
	}

	slog.DebugCtx(ctx, "track", track)

	return track, nil
}
