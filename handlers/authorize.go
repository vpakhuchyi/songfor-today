package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slog"

	"github.com/vpakhuchyi/songfor-today/adapters/auth"
	"github.com/vpakhuchyi/songfor-today/config"
	"github.com/vpakhuchyi/songfor-today/models"
	"github.com/vpakhuchyi/songfor-today/secrets"
)

const deezerAccessToken = "deezer-access-token"

type Authorizer struct {
	auth   auth.Client
	config config.Config
	sm     *secrets.Client
}

func NewAuthorizer(auth auth.Client, cfg config.Config, sm *secrets.Client) Authorizer {
	return Authorizer{
		auth:   auth,
		config: cfg,
		sm:     sm,
	}
}

func (a *Authorizer) Authorize(ctx *fiber.Ctx) error {
	redirectURI := fmt.Sprintf(
		"%s?redirect_uri=%s&app_id=%s&perms=%s",
		a.config.Authorization.LoginURL,
		a.config.Authorization.RedirectURI,
		a.config.Authorization.AppID,
		a.config.Authorization.Permissions,
	)

	slog.DebugCtx(ctx.Context(), "Making an auth request", "redirect_uri", redirectURI)

	if err := ctx.Redirect(redirectURI, http.StatusMovedPermanently); err != nil {
		return fmt.Errorf("failed to redirect to Deezer server: %w", err)
	}

	return nil
}

func (a *Authorizer) GetAccessToken(ctx *fiber.Ctx) error {
	p := models.GetAccessTokenParams{
		AppID:  a.config.Authorization.AppID,
		Secret: a.config.DeezerAppSecret,
		Code:   ctx.Query("code"),
	}

	token, err := a.auth.GetAccessToken(ctx.Context(), p)
	if err != nil {
		slog.Error("Failed to complete the login", "err", err)
		return err
	}

	if err := a.sm.Store(ctx.Context(), deezerAccessToken, token); err != nil {
		slog.Error("Failed to store Deezer access-token", "err", err)
		return err
	}

	return nil
}
