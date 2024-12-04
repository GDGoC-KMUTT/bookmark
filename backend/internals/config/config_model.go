package config

type EnvModel struct {
	DBAutoMigrate bool     `json:"DB_AUTOMIGRATE" mapstructure:"DB_AUTOMIGRATE"`
	DBHost        string   `json:"DB_HOST" mapstructure:"DB_HOST"`
	DBName        string   `json:"DB_NAME" mapstructure:"DB_NAME"`
	DBPassword    string   `json:"DB_PASSWORD" mapstructure:"DB_PASSWORD"`
	DBPort        int      `json:"DB_PORT" mapstructure:"DB_PORT"`
	DBUsername    string   `json:"DB_USERNAME" mapstructure:"DB_USERNAME"`
	ServerHost    string   `json:"SERVER_HOST" mapstructure:"SERVER_HOST"`
	ServerOrigins []string `json:"SERVER_ORIGINS" mapstructure:"SERVER_ORIGINS"`
	ServerPort    int      `json:"SERVER_PORT" mapstructure:"SERVER_PORT"`
	SecretKey     string   `json:"SECRET" mapstructure:"SECRET"`
	Environment   int      `json:"ENVIRONMENT" mapstructure:"ENVIRONMENT"`
}
