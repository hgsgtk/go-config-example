package config_test

import (
	"testing"

	"os"

	"github.com/Khigashiguchi/go-config-example/src/config"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cases := []struct {
		name     string
		appMode  string
		expected config.Config
	}{
		{
			name:    "localhost",
			appMode: "localhost",
			expected: config.Config{
				DB: config.DBConfig{
					User:     "sample_user",
					Password: "sample_password",
					Host:     "master_mysql",
					Port:     3306,
					Name:     "sample",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			confDir := "./env/"
			os.Setenv("DB_PASSWORD", c.expected.DB.Password)

			res, err := config.NewConfig(confDir, c.appMode)

			assert.Equal(t, nil, err)
			assert.Equal(t, c.expected.DB.User, res.DB.User)
			assert.Equal(t, c.expected.DB.Password, res.DB.Password)
			assert.Equal(t, c.expected.DB.Host, res.DB.Host)
			assert.Equal(t, c.expected.DB.Port, res.DB.Port)
			assert.Equal(t, c.expected.DB.Name, res.DB.Name)
		})
	}
}
