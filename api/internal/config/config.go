package config

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	Port          int
	PostgreSQLURL *url.URL
}

func (c *Config) LoadFromEnv() error {
	if p, found := os.LookupEnv("PORT"); found {
		port, err := strconv.Atoi(p)
		if err != nil {
			return fmt.Errorf("port must be numerical value: %w", err)
		}
		c.Port = port
	}

	if pURL, found := os.LookupEnv("POSTGRESQL_URL"); found {
		postgresURL, err := url.Parse(pURL)
		if err != nil {
			return fmt.Errorf("invalid postgres url: %w", err)
		}
		c.PostgreSQLURL = postgresURL

	}
	return nil
}

func (c *Config) LoadDefault() {
	c.Port = 53430
}

func (c *Config) Load() error {
	c.LoadDefault()
	if err := c.LoadFromEnv(); err != nil {
		return fmt.Errorf("unable to load from env: %w", err)
	}
	return c.Validate()
}

func (c *Config) Validate() error {
	if c.Port == 0 {
		return fmt.Errorf("port is required")
	}

	if c.PostgreSQLURL == nil {
		return fmt.Errorf("postgres url is required")
	}

	return nil
}
