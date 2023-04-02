package auth

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/vpakhuchyi/songfor-today/config"
	"github.com/vpakhuchyi/songfor-today/models"
)

const (
	secret = "secret"
	code   = "code"
)

type Client struct {
	httpclient *resty.Client
	cfg        config.Authorization
}

func New(httpclient *resty.Client, cfg config.Authorization) Client {
	return Client{
		httpclient: httpclient,
		cfg:        cfg,
	}
}

func (c *Client) GetAccessToken(ctx context.Context, params models.GetAccessTokenParams) (string, error) {
	resp, err := c.httpclient.SetDebug(true).
		R().
		SetContext(ctx).
		SetQueryParams(
			map[string]string{
				secret: params.Secret,
				code:   params.Code,
			},
		).
		Get(c.cfg.TokenURL)
	if err != nil {
		return "", fmt.Errorf("failed to reach Deezer server: %w", err)
	}

	return resp.String(), nil
}
