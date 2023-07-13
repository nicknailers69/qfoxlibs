package network

import (
	"context"
	"crypto/tls"
	"sync"

	"go.uber.org/zap"
)

type Server struct {
	conf      *NetworkConfig
	tls       *tls.Config
	clientTls *tls.Config

	ctx    context.Context
	cancel context.CancelFunc

	pool *sync.Pool

	mu     sync.Mutex
	wg     sync.WaitGroup
	log    *zap.Logger
	kemPub []byte
}
