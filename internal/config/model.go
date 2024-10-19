package config

type Config struct {
	Server   Server
	Database Database
	Jwt      Jwt
}

type Jwt struct {
	Key string
	Exp int
}

type Server struct {
	Host string
	Port string
}

type Database struct {
	Host string
	Port string
	Name string
	User string
	Pass string
	Tz   string
}
