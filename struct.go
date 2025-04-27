package main

type Config struct {
	ServerName string `json:"servername"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Database   string `json:"database"`
}

type Params struct {
	FHD string `json:"FHD"`
	KHD string `json:"KHD"`
}

type AppConfig struct {
	Config Config `json:"config"`
	Params Params `json:"params"`
}
