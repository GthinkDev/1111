package config

import "fmt"

type System struct {
	Env  string `yaml:"env"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func (s System) Addr() string {
	return fmt.Sprintf("%s:%s", s.Host, s.Port)
}
