/*
 * Copyright © 2021-2022 Durudex

 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.

 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.

 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type (
	Config struct {
		Server  ServerConfig
		Auth    AuthConfig
		Service ServiceConfig
	}

	// Server config variables.
	ServerConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		Name string `mapstructure:"name"`
	}

	// Auth config variables.
	AuthConfig struct {
		JWT JWTConfig
	}

	// JWT config variables.
	JWTConfig struct {
		SigningKey string
	}

	// Service base config.
	Service struct {
		Addr string `mapstructure:"addr"`
		TLS  bool   `mapstructure:"tls"`
	}

	// Services config varibles.
	ServiceConfig struct {
		Auth Service
		Code Service
	}
)

// Initialize config.
func Init() (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Populate defaults config variables.
	populateDefaults()

	// Parsing config file.
	if err := parseConfigFile(); err != nil {
		return nil, fmt.Errorf("error parsing config file: %s", err.Error())
	}

	var cfg Config
	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshal config keys: %s", err.Error())
	}

	// Set env configurations.
	setFromEnv(&cfg)

	return &cfg, nil
}

// Parsing config file.
func parseConfigFile() error {
	// Get config path variable.
	configPath := os.Getenv("CONFIG_PATH")

	// Check is config path variable empty.
	if configPath == "" {
		configPath = defaultConfigPath
	}

	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(path[0]) // Folder.
	viper.SetConfigName(path[1]) // File.

	// Read config file.
	return viper.ReadInConfig()
}

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal server keys.
	if err := viper.UnmarshalKey("server", &cfg.Server); err != nil {
		return err
	}
	// Unmarshal code service keys.
	if err := viper.UnmarshalKey("service.code", &cfg.Service.Code); err != nil {
		return err
	}
	// Unmarshal auth service keys.
	return viper.UnmarshalKey("service.auth", &cfg.Service.Auth)
}

// Seting environment variables from .env file.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set from environment configurations...")

	// Auth variables.
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables.")

	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("server.name", defaultServerName)
}
