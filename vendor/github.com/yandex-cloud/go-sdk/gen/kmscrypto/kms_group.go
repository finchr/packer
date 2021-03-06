// Code generated by sdkgen. DO NOT EDIT.

package kms

import (
	"context"

	"google.golang.org/grpc"
)

// KMSCrypto provides access to "kms" component of Yandex.Cloud
type KMSCrypto struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

// NewKMSCrypto creates instance of KMSCrypto
func NewKMSCrypto(g func(ctx context.Context) (*grpc.ClientConn, error)) *KMSCrypto {
	return &KMSCrypto{g}
}

// SymmetricCrypto gets SymmetricCryptoService client
func (k *KMSCrypto) SymmetricCrypto() *SymmetricCryptoServiceClient {
	return &SymmetricCryptoServiceClient{getConn: k.getConn}
}
