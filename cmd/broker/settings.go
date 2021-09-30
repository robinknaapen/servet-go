package main

type Database struct {
	ConnectionString string `env:"DATABASE_CONNECTION_STRING,required"`
}

var settings = struct {
	Database Database
}{}
