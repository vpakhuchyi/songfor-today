package secrets

import (
	"context"
	"fmt"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
	"golang.org/x/exp/slog"
)

type Client struct {
	sm        *secretmanager.Client
	projectID string
}

func NewClient(ctx context.Context, projectID string) (*Client, error) {
	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize secret manager: %w", err)
	}

	slog.InfoCtx(ctx, "Successfully connected to GCP Secrets Manager")

	return &Client{sm: client, projectID: projectID}, nil
}

func (c *Client) Store(ctx context.Context, name string, value string) error {
	slog.DebugCtx(ctx, "Request to store a secret", "name", name)

	createSecretReq := &secretmanagerpb.CreateSecretRequest{
		Parent:   fmt.Sprintf("projects/%s", c.projectID),
		SecretId: name,
		Secret: &secretmanagerpb.Secret{
			Replication: &secretmanagerpb.Replication{
				Replication: &secretmanagerpb.Replication_Automatic_{
					Automatic: &secretmanagerpb.Replication_Automatic{},
				},
			},
		},
	}

	secret, err := c.sm.CreateSecret(ctx, createSecretReq)
	if err != nil {
		return fmt.Errorf("failed to create secret: %w", err)
	}

	addSecretVersionReq := &secretmanagerpb.AddSecretVersionRequest{
		Parent: secret.Name,
		Payload: &secretmanagerpb.SecretPayload{
			Data: []byte(value),
		},
	}

	_, err = c.sm.AddSecretVersion(ctx, addSecretVersionReq)
	if err != nil {
		return fmt.Errorf("failed to add secret version: %w", err)
	}

	return nil
}

func (c *Client) Load(ctx context.Context, name string) (string, error) {
	slog.DebugCtx(ctx, "Request to load a secret", "name", name)

	accessRequest := &secretmanagerpb.AccessSecretVersionRequest{
		Name: name,
	}

	result, err := c.sm.AccessSecretVersion(ctx, accessRequest)
	if err != nil {
		return "", fmt.Errorf("failed to access secret version: %w", err)
	}

	return result.Payload.String(), nil
}

func (c *Client) IsExists(ctx context.Context, name string) (bool, error) {
	slog.DebugCtx(ctx, "Request to get a secret metadata", "name", name)

	getRequest := &secretmanagerpb.GetSecretRequest{
		Name: name,
	}

	_, err := c.sm.GetSecret(ctx, getRequest)
	if err != nil {
		if err.Error() == "rpc error: code = NotFound desc = Resource not found" {
			return false, nil
		}

		return false, fmt.Errorf("failed to get secret metadata: %w", err)
	}

	return true, nil
}
