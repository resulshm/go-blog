package config

type App struct {
	Name string
}

type Server struct {
	Host string
	Port string
}

type Config struct {
	App    App
	Server Server
	DB     DB
}

type DB struct {
	Username string
	Password string
	Host     string
	Port     string
	Name     string
}
