// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import (
	botconfig "github.com/mxyns/justlog/config"
	"log"
	"strings"
	"time"
)

type Config struct {
	Period             time.Duration   `config:"period"`
	StreamStatusPeriod time.Duration   `config:"streamStatusCheckPeriod"`
	StreamStatusOffset time.Duration   `config:"streamStatusCheckOffset"`
	QueueCapacity      uint64          `config:"queueCapacity"`
	BotVerified        bool            `config:"botVerified"`
	AdminAPIKey        string          `config:"adminAPIKey"`
	Username           string          `config:"username"`
	Oauth              string          `config:"oauth"`
	ListenAddress      string          `config:"listenAddress"`
	Admins             []string        `config:"admins"`
	Channels           []string        `config:"channels"`
	ClientID           string          `config:"clientID"`
	ClientSecret       string          `config:"clientSecret"`
	LogLevel           string          `config:"logLevel"`
	OptOut             map[string]bool `config:"optOut"`
}

var DefaultConfig = Config{
	Period:             16 * time.Second,
	StreamStatusPeriod: 16 * time.Second,
	StreamStatusOffset: 8 * time.Second,
	QueueCapacity:      1024,
	BotVerified:        true,
	AdminAPIKey:        "noshot",
	Username:           "Username",
	Oauth:              "OAuth",
	ListenAddress:      ":8025",
	Admins:             []string{},
	Channels:           []string{},
	ClientID:           "ClientId",
	ClientSecret:       "ClientSecret",
	LogLevel:           "info",
	OptOut:             map[string]bool{},
}

func ConvertConfiguration(config *Config) *botconfig.Config {
	// setup defaults
	cfg := botconfig.Config{
		LogsDirectory: "./logs",
		ListenAddress: ":8025",
		Username:      "justinfan777777",
		OAuth:         "oauth:777777777",
		Channels:      []string{},
		Admins:        []string{"gempir"},
		LogLevel:      "info",
		Archive:       false,
		OptOut:        map[string]bool{},
	}

	cfg.ClientSecret = config.ClientSecret
	cfg.Channels = config.Channels
	cfg.OptOut = config.OptOut
	cfg.LogLevel = config.LogLevel
	cfg.Admins = config.Admins
	cfg.ListenAddress = config.ListenAddress
	cfg.Username = config.Username
	cfg.AdminAPIKey = config.AdminAPIKey
	cfg.BotVerified = config.BotVerified
	cfg.ClientID = config.ClientID
	cfg.OAuth = config.Oauth

	// normalize
	cfg.LogsDirectory = strings.TrimSuffix(cfg.LogsDirectory, "/")
	cfg.OAuth = strings.TrimPrefix(cfg.OAuth, "oauth:")
	cfg.LogLevel = strings.ToLower(cfg.LogLevel)
	cfg.SetupLogger()

	// ensure required
	if cfg.ClientID == "" {
		log.Fatal("No clientID specified")
	}

	return &cfg
}
