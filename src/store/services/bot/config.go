package bot

type Config struct {
	Token string
}

func NewConfig(token string) *Config {
	config := Config{}
	config.Token = token
	return &config
}

func NewDefaultConfig() *Config{
	return NewConfig("426666089:AAE7mt0918YCmA1bnJMi8ORKT9at3Y_pLEk")
}
