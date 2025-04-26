package main

type Config struct {
	ServerName string `json:"servername"`
	Port       int    `json:"port"`
	User       string `json:"user"`
	Password   string `json:"password"`
	Database   string `json:"database"`
}

type Params struct {
	FirstTable  string `json:"first_table"`
	SecondTable string `json:"second_table"`
}

type AppConfig struct {
	Config Config `json:"config"`
	Params Params `json:"params"`
}
