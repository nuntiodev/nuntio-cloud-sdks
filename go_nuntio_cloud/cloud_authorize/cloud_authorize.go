package cloud_authorize

import (
	"context"
)

type CloudAuthorize interface {
	GetAccessToken(ctx context.Context) (string, error)
}

type noAuthorization struct{}

func (a *noAuthorization) GetAccessToken(ctx context.Context) (string, error) {
	return "", nil
}

func New(authorize CloudAuthorize) (CloudAuthorize, error) {
	if authorize != nil {
		return authorize, nil
	}
	return &noAuthorization{}, nil
}
