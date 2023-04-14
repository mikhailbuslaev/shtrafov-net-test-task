package cmd

import (
	"github.com/spf13/pflag"
)

// Config ...
type Config struct {
	TcpAddr string
}

// Prepare ...
func (c *Config) Prepare() (err error) {
	return nil
}

func (c *Config) Flags() *pflag.FlagSet {
	f := pflag.NewFlagSet("", pflag.PanicOnError)
	// ENV TCPADDR
	f.StringVar(&c.TcpAddr, "port", "80", "port for server")

	return f
}

func (c *Config) Validate() error {
	return nil
}
