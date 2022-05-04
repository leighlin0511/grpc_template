package config

import "time"

type Configuration struct {
	Server ServerConfig
}

type ServerConfig struct {
	GrpcPort        int
	HTTPPort        int
	ShutdownTimeout time.Duration
}
