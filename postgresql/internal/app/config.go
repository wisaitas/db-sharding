package app

var Config struct {
	Server struct {
		Port string `env:"PORT" envDefault:"8080"`
	} `envPrefix:"SERVER_"`
}
