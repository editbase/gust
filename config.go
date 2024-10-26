// config.go
package gust

type Config struct {
	Port        string
	TemplateDir string
	StaticDir   string
	DevMode     bool
}

func defaultConfig() *Config {
	return &Config{
		Port:        "3000",
		TemplateDir: "./templates",
		StaticDir:   "./static",
		DevMode:     false,
	}
}
