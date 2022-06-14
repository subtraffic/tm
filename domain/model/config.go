package vb

type Config struct {
	N int
	C int
	H string
}

var configIns *Config

func NewConfig() *Config {
	if configIns == nil {
		configIns = &Config{
			N: 10000,
			C: 100,
			H: "http://127.0.0.1:8000",
		}
	}
	return configIns
}

func init() {
	config := NewConfig()
	config.C = 100
	config.N = 10000
	config.H = "http://127.0.0.1"
}
