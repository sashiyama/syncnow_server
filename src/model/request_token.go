package model

import (
	"errors"
	"strings"
)

type RequestToken struct {
	Token string
}

func GetRequestTokenFromHeader(header map[string][]string) (RequestToken, error) {
	authorizations, ok := header["Authorization"]

	if !ok {
		return RequestToken{}, errors.New("Authorization header is missing")
	}

	const prefix = "Bearer "

	if !strings.HasPrefix(authorizations[0], prefix) {
		return RequestToken{}, errors.New("Bearer token is missing")
	}

	return RequestToken{Token: strings.TrimPrefix(authorizations[0], prefix)}, nil
}
