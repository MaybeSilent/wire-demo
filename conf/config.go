package conf

type Config struct {
	conf string
}

func (c *Config) GetConf() string {
	return c.conf
}

func InitConfig() *Config {
	return &Config{conf: "settings"}
}
