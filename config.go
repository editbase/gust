// /config.go
// Config holds the application configuration settings

package stardust

// Config holds the application configuration settings
type Config struct {
	Port        string // HTTP server port
	TemplateDir string // Directory containing HTML templates
	StaticDir   string // Directory containing static files
	DevMode     bool   // Enables development mode features
}

// defaultConfig returns a default configuration object
func defaultConfig() *Config {
	return &Config{
		Port:        "3000",
		TemplateDir: "./templates",
		StaticDir:   "./static",
		DevMode:     false,
	}
}
