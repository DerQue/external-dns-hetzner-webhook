package configuration

import (
	"github.com/caarlos0/env/v8"
	log "github.com/sirupsen/logrus"
)

// Config struct for configuration environmental variables
type Config struct {
	ServerHost string `env:"SERVER_HOST" envDefault:"localhost"`
	ServerPort int    `env:"SERVER_PORT" envDefault:"8888"`

	DomainFilter         []string `env:"DOMAIN_FILTER" envDefault:""`
	ExcludeDomains       []string `env:"EXCLUDE_DOMAIN_FILTER" envDefault:""`
	RegexDomainFilter    string   `env:"REGEXP_DOMAIN_FILTER" envDefault:""`
	RegexDomainExclusion string   `env:"REGEXP_DOMAIN_FILTER_EXCLUSION" envDefault:""`
}

// Init sets up configuration by reading set environmental variables
func Init() Config {
	cfg := Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Error reading configuration from environment: %v", err)
	}
	return cfg
}
