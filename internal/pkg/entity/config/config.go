package config

// Config declare configuration of application
type Config struct {
	APIKey    string `json:"api_key" mapstructure:"api_key"`
	Base      string `json:"base" mapstructure:"base"`
	StartDate string `json:"start_date" mapstructure:"start_date"`
}
