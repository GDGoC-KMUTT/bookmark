package config

type Config struct {
	DBAutoMigrate     *bool     `yaml:"DB_AUTOMIGRATE" mapstructure:"DB_AUTOMIGRATE"`
	DBHost            *string   `yaml:"DB_HOST" mapstructure:"DB_HOST"`
	DBName            *string   `yaml:"DB_NAME" mapstructure:"DB_NAME"`
	DBPassword        *string   `yaml:"DB_PASSWORD" mapstructure:"DB_PASSWORD"`
	DBPort            *int      `yaml:"DB_PORT" mapstructure:"DB_PORT"`
	DBUsername        *string   `yaml:"DB_USERNAME" mapstructure:"DB_USERNAME"`
	ServerHost        *string   `yaml:"SERVER_HOST" mapstructure:"SERVER_HOST"`
	ServerOrigins     []*string `yaml:"SERVER_ORIGINS" mapstructure:"SERVER_ORIGINS"`
	ServerPort        *int      `yaml:"SERVER_PORT" mapstructure:"SERVER_PORT"`
	SecretKey         *string   `yaml:"SECRET" mapstructure:"SECRET"`
	Environment       *int      `yaml:"ENVIRONMENT" mapstructure:"ENVIRONMENT"`
	OauthClientId     *string   `yaml:"OAUTH_CLIENT_ID" mapstructure:"OAUTH_CLIENT_ID"`
	OauthClientSecret *string   `yaml:"OAUTH_CLIENT_SECRET" mapstructure:"OAUTH_CLIENT_SECRET"`
	OauthEndpoint     *string   `yaml:"OAUTH_ENDPOINT" mapstructure:"OAUTH_ENDPOINT"`
	FrontendUrl       *string   `yaml:"FRONTEND_URL" mapstructure:"FRONTEND_URL"`
	FrontendScheme    *string   `yaml:"FRONTEND_SCHEME" mapstructure:"FRONTEND_SCHEME"`
}
