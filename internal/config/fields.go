package config

type Service struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

type Redis struct {
	Address string `yaml:"address"`
	DB      int    `yaml:"db"`
}
