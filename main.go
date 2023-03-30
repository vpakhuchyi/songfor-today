package main

import (
	"context"

	gcpdatastore "cloud.google.com/go/datastore"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"

	"github.com/vpakhuchyi/songfor-today/adapters/auth"
	"github.com/vpakhuchyi/songfor-today/adapters/datastore"
	"github.com/vpakhuchyi/songfor-today/config"
	"github.com/vpakhuchyi/songfor-today/handlers"
	"github.com/vpakhuchyi/songfor-today/logger"
	"github.com/vpakhuchyi/songfor-today/secrets"
	"github.com/vpakhuchyi/songfor-today/usecases"
)

func main() {
	app := fiber.New()
	cfg := config.MustReadConfiguration()

	logger.SetLevel(cfg.LogLevel)
	ctx := context.Background()

	sm, err := secrets.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		slog.ErrorCtx(ctx, "Failed to create secrets manager client", "err", err)

		return
	}

	ds, err := gcpdatastore.NewClient(ctx, cfg.ProjectID)
	if err != nil {
		slog.ErrorCtx(ctx, "Failed to create client", "err", err)

		return
	}
	defer ds.Close()

	datastoreAdapter := datastore.New(ds)
	restyClient := resty.New()
	authAdapter := auth.New(restyClient, cfg.Authorization)
	authClient := handlers.NewAuthorizer(authAdapter, cfg, sm)
	//deezerAdapter := deezer.New(restyClient)
	tracksUsecases := usecases.New(&datastoreAdapter)
	tracksClient := handlers.NewTracks(tracksUsecases, cfg)

	app.Static("/termsofuse", "./static/termsofuse.html")

	app.Get("/auth", authClient.Authorize)
	app.Get("/deezercallback", authClient.GetAccessToken)
	app.Get("/", tracksClient.RandomSong)

	if err := app.Listen(":8080"); err != nil {
		slog.ErrorCtx(ctx, "Shutting down the app", "err", err)

		app.Shutdown()
	}
}
