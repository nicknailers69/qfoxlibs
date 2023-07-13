package network

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type NetConf interface {
	GetEncoder() *toml.Encoder
	GetDecoder() *toml.Decoder
}

type NetworkConfig struct {
	Cfg    *GeneralConfig
	Listen []ListenerConfig
}

type GeneralConfig struct {
	Logging  string `toml:"log,false,false"`
	LogLevel string `toml:"logLevel,false,false"`
	Uid      string `toml:"uid"`
	Gid      string `toml:"gid"`
	ConfPath string `toml:"configPath"`
}

type ListenerConfig struct {
	Addr      string
	NodeID    string
	Allow     []subnet
	Deny      []subnet
	Timeout   Timeouts
	Quic      bool
	Tls       *TlsServerConfig
	RateLimit *RateLimit
	Connect   ConnectConfig
	serverCfg *tls.Config
	clientCfg *tls.Config
}

type Timeouts struct {
	Connect int
	Read    int
	Write   int
}

type RateLimit struct {
	Global    int
	PerHost   int
	CacheSize int
}

type subnet struct {
	net.IPNet
}

type ConnectConfig struct {
	Addr          string
	Bind          string
	ProxyProtocol string
	Quic          bool
	Tls           *TlsClientConfig
}

type TlsServerConfig struct {
	Quic       bool
	Sni        string
	Cert       string
	Key        string
	ClientCert string
	Server     string
	ClientCA   string
}

type TlsClientConfig struct {
	Quic   bool
	Ca     string
	Cert   string
	Key    string
	Server string
	tlsCfg *tls.Config
}

func NewConfig(fn string) (*NetworkConfig, error) {

	data, err := os.ReadFile(fn)
	if err != nil {
		return nil, fmt.Errorf("cannot read config file: %s", fn)
	}
	var cfg NetworkConfig
	err = toml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("cannot parse config file: %s", fn)
	}
	return NetworkConfigDefaults(&cfg), nil
}

func NetworkConfigDefaults(nc *NetworkConfig) *NetworkConfig {
	for _, l := range nc.Listen {

		l.RateLimit = &RateLimit{1000, 10, 5000}

		t := &l.Timeout
		if t.Connect == 0 {
			t.Connect = 5
		}
		if t.Read == 0 {
			t.Read = 2
		}

		if t.Write == 0 {
			t.Write = 2
		}

	}
	if len(nc.Cfg.LogLevel) == 0 {
		nc.Cfg.LogLevel = "INFO"
	}

	if len(nc.Cfg.Logging) == 0 {
		nc.Cfg.Logging = "SYSLOG"
	}
	return nc
}
