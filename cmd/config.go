package cmd

import (
	"github.com/spf13/pflag"
)

// Config ...
type Config struct {
	TcpAddr  string
	HttpPort string
}

// Prepare ...
func (c *Config) Prepare() (err error) {
	return nil
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("", pflag.PanicOnError)
	// ENV TCPADDR
	f.StringVar(&c.TcpAddr, "tcpAddr", "", "tcp address for server")
	// ENV HTTPPORT
	f.StringVar(&c.HttpPort, "httpPort", "", "port for http interface")

	return f
}

func (c *Config) Validate() error {
	return nil
}
