package config

// DBConfig represents database connection configuration information.
type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Name     string
}

// Config represents application configuration.
type Config struct {
	DB DBConfig
}

// NewConfig return configuration struct.
func NewConfig() (Config, error) {
	conf := Config{}

	conf.DB = DBConfig{
		User:     "sample_user",
		Password: "sample_password",
		Host:     "master_mysql",
		Port:     3306,
		Name:     "sample",
	}

	return conf, nil
}
