package config

import "os"

type NSQd struct {
	Host     string
	TCPPort  string
	HTTPPort string
}

type Config struct {
	NSQd
}

func Parse() Config {
	nsqd := NSQd{
		Host:     os.Getenv("NSQD_HOST"),
		TCPPort:  os.Getenv("NSQD_TCP_PORT"),
		HTTPPort: os.Getenv("NSQD_HTTP_PORT"),
	}

	return Config{
		nsqd,
	}
}
