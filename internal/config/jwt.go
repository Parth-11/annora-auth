package config

import (
	"time"
)

type JWTConfig struct {
	PrivateKeyPath string
	PublicKeyPath  string
	Issuer     string
	AccessTTL  time.Duration
	RefreshTTL time.Duration
}