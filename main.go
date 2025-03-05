package main

import (
	"encoding/json"
	"io/ioutil"
	"jsondiff/diffServer"
	"os"
)

type ConfigServer struct {
	Header    map[string]string `json:"header"`
	Parameter map[string]string `json:"parameter"`
}

func (e *ConfigServer) Init() {
	e.Header = map[string]string{
		"token": "",
		"type":  "",
	}
	e.Parameter = map[string]string{
		"id":   "",
		"user": "",
	}
}

type Config struct {
	DataServerA  ConfigServer `json:"dataServerA"`
	TokenServerA ConfigServer `json:"tokenServerA"`
	DataServerB  ConfigServer `json:"dataServerB"`
	TokenServerB ConfigServer `json:"tokenServerB"`
}

func (e *Config) Init() {
	e.DataServerA.Init()
	e.TokenServerA.Init()
	e.DataServerB.Init()
	e.TokenServerB.Init()
}

func main() {
	// Detecta se o arquivo de configuração existe e carrega o arquivo
	if _, err := os.Stat("config.json"); err == nil {
		file, _ := ioutil.ReadFile("config.json")
		config := new(Config)
		_ = json.Unmarshal(file, &config)
	} else {
		config := new(Config)
		config.Init()
	}
	c := new(diffServer.Console)
	c.Init()
}
