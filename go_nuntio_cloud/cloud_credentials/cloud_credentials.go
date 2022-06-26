package cloud_credentials

import (
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type CloudCredentials interface {
	GetTransportCredentials() (credentials.TransportCredentials, error)
}
type insecureTransportCredentials struct{}

func (tc *insecureTransportCredentials) GetTransportCredentials() (credentials.TransportCredentials, error) {
	return insecure.NewCredentials(), nil
}

func New(transportCredentials CloudCredentials) (CloudCredentials, error) {
	if transportCredentials != nil {
		return transportCredentials, nil
	}
	return &insecureTransportCredentials{}, nil
}
