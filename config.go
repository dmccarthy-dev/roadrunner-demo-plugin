package roadrunner_demo_plugin

type Config struct {
	Message string `mapstructure:"message"`
}

// InitDefaults .. You can also initialize some defaults values for config keys
func (cfg *Config) InitDefaults() {
	if cfg.Message == "" {
		cfg.Message = "default hello world"
	}
}
