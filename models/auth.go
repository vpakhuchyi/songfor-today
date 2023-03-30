package models

type LoginParams struct {
	AppID       string
	Permissions string
	RedirectURI string
}

type GetAccessTokenParams struct {
	AppID  string
	Code   string
	Secret string
}
