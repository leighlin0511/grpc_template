package config

type Configuration struct {
	Server ServerConfig
}

type ServerConfig struct {
	GrpcPort int
	HTTPPort int
}
