package config

import "strconv"

// Mysql ���置

type Mysql struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Charset  string `yaml:"charset"`
	LogLevel string `yaml:"log_level"`
	Config   string `yaml:"config"`
}

func (m Mysql) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + strconv.Itoa(m.Port) + ")/" + m.Database + "?charset=" + m.Charset + "&parseTime=True&loc=Local" + "&" + m.Config
}
