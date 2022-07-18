/*
 * Copyright © 2021-2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package config

import (
	"os"
	"path/filepath"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default config path.
const defaultConfigPath string = "configs/main"

type (
	// Config variables.
	Config struct {
		HTTP    HTTPConfig
		Service ServiceConfig
		Auth    AuthConfig
	}

	// HTTP server config variables.
	HTTPConfig struct {
		Host string `mapstructure:"host"`
		Port string `mapstructure:"port"`
		Name string `mapstructure:"name"`
	}

	// TLS config variables.
	TLSConfig struct {
		Enable bool   `mapstructure:"enable"`
		CACert string `mapstructure:"ca-cert"`
		Cert   string `mapstructure:"cert"`
		Key    string `mapstructure:"key"`
	}

	// Auth config variables.
	AuthConfig struct {
		JWT JWTConfig `mapstructure:"jwt"`
	}

	// JWT config variables.
	JWTConfig struct{ SigningKey string }

	// Service base config.
	Service struct {
		Addr string    `mapstructure:"addr"`
		TLS  TLSConfig `mapstructure:"tls"`
	}

	// Services config variables.
	ServiceConfig struct {
		User Service `mapstructure:"user"`
		Post Service `mapstructure:"post"`
	}
)

// Initialize config.
func Init() (*Config, error) {
	log.Debug().Msg("Initialize config...")

	// Parsing config file.
	if err := parseConfigFile(); err != nil {
		return nil, err
	}

	var cfg Config
	// Unmarshal config keys.
	if err := unmarshal(&cfg); err != nil {
		return nil, err
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
	dir, file := filepath.Split(configPath)

	viper.AddConfigPath(dir)
	viper.SetConfigName(file)

	// Read config file.
	return viper.ReadInConfig()
}

// Unmarshal config keys.
func unmarshal(cfg *Config) error {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal server keys.
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		return err
	}
	// Unmarshal service keys.
	return viper.UnmarshalKey("service", &cfg.Service)
}

// Setting environment variables from .env file.
func setFromEnv(cfg *Config) {
	log.Debug().Msg("Set from environment configurations...")

	// Auth variables.
	cfg.Auth.JWT.SigningKey = os.Getenv("JWT_SIGNING_KEY")
}
