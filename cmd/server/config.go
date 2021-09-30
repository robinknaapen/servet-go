package main

type Database struct {
	ConnectionString string `env:"DATABASE_CONNECTION_STRING,required"`
}

type Cors struct {
	AllowOrigin []string `env:"CORS_ALLOW_ORIGIN" envSeparator:","`
}

var settings = struct {
	Database Database
	Cors     Cors
}{}
