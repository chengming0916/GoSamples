package config

type Conifg struct {
	Port       uint   `yaml:"Port"`
	DbHost     string `yaml:"DbHost"`
	DbPort     uint   `yaml:"DbPort"`
	DbName     string `yaml:"DbName"`
	DbUser     string `yaml:"DbUser"`
	DbPassword string `yaml:"DbPassword"`
}

// func ReadConfig() *Config {

// }
