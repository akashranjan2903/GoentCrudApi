package config

type Config struct {
	Database struct {
		Dsn string `mapstructure:"dsn"`
	} `mapstructure:"database"`
	App struct {
		Port string `mapstructure:"port"`
		Env  string `mapstructure:"env"`
	} `mapstructure:"app"`
	Pagination struct {
		Offset int `mapstructure:"offset"`
		Limit  int `mapstructure:"limit"`
	} `mapstructure:"pagination"`
	JWT struct {
		Secret           string `mapstructure:"secret"`
		ExpiryMinAccess  int    `mapstructure:"expiry"`
		ExpiryMinRefresh int    `mapstructure:"expiry_refresh"`
		ExpiryMinReset   int    `mapstructure:"expiry_reset"`
	} `mapstructure:"jwt"`
}

func GetConfig() Config {
	// Change the env variable to the one you want to use
	return InitConfig(Environment(Dev))
}
