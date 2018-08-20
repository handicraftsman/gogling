package main

import (
  "log"
  "encoding/json"
  "io/ioutil"
  "os"
)

// Config - stores gogling configuration
type Config struct {
	Host           string
	Port           string
	RootDir        string
	RouterFile     string
  ServeTLS       bool
}

// ConfigLoad - loads json config into a Config struct instance
func ConfigLoad(path string) *Config {
  l := log.New(StdoutFile, "config ", -1)
  l.Printf("Reading config from `%s`", path)

  cfg := &Config{}

  data, err := ioutil.ReadFile(path)
  if err != nil {
    l.Fatal(err)
  }

  err = json.Unmarshal(data, cfg)
  if err != nil {
    l.Fatal(err)
  }

  if cfg.Host == "ENV" {
    cfg.Host = os.Getenv("HOST")
  }

  if cfg.Port == "ENV" {
    cfg.Port = os.Getenv("PORT")
  }

  l.Printf("Done")

  return cfg
}