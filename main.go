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
	e.Header = make(map[string]string)
	e.Parameter = make(map[string]string)
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
	config := new(Config)
	if _, err := os.Stat("config.json"); err == nil {
		file, _ := ioutil.ReadFile("config.json")
		_ = json.Unmarshal(file, &config)
	} else {
		config.Init()
		file, _ := json.MarshalIndent(config, "", " ")
		_ = ioutil.WriteFile("config.json", file, 0644)
	}
	c := new(diffServer.Console)
	for k, v := range config.DataServerA.Header {
		c.AddHeaderServerA(k, v)
	}
	for k, v := range config.DataServerA.Parameter {
		c.AddParamServerA(k, v)
	}
	for k, v := range config.DataServerB.Header {
		c.AddHeaderServerB(k, v)
	}
	for k, v := range config.DataServerB.Parameter {
		c.AddParamServerB(k, v)
	}
	c.Init()
}
