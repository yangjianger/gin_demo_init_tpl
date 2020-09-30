package config

type Config struct {
	*MySQLConfig `mapstructure:"mysql"`
	*RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}
