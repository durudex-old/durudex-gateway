/*
	Copyright © 2021 Durudex

	This file is part of Durudex: you can redistribute it and/or modify
	it under the terms of the GNU Affero General Public License as
	published by the Free Software Foundation, either version 3 of the
	License, or (at your option) any later version.

	Durudex is distributed in the hope that it will be useful,
	but WITHOUT ANY WARRANTY; without even the implied warranty of
	MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
	GNU Affero General Public License for more details.

	You should have received a copy of the GNU Affero General Public License
	along with Durudex. If not, see <https://www.gnu.org/licenses/>.
*/

package config

import (
	"os"
	"strings"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

// Default variables.
const (
	defaultHTTPPort    = "8000"
	defaultHTTPAppName = "durudex-gateway"
)

type (
	Config struct {
		HTTP    HTTPConfig
		GRPC    GRPCConfig
		Service ServiceConfig
		Auth    AuthConfig
	}

	// HTTP config variables.
	HTTPConfig struct {
		Host    string `mapstructure:"host"`
		Port    string `mapstructure:"port"`
		AppName string `mapstructure:"appName"`
	}

	// GRPC config variables.
	GRPCConfig struct {
		TLS bool `mapstructure:"tls"`
	}

	// Auth config variables.
	AuthConfig struct {
		// Signing auth key.
		SigningKey string
	}

	// Service config varibles.
	ServiceConfig struct {
		// Auth service variables.
		Auth AuthServiceConfig
	}

	// Auth service config variables.
	AuthServiceConfig struct {
		Addr string `mapstructure:"addr"`
	}
)

// Initialize config.
func Init(configPath string) *Config {
	log.Debug().Msg("Initialize config...")

	// Parsing config file.
	parseConfigFile(configPath)

	var cfg Config
	// Unmarshal config keys.
	unmarshal(&cfg)

	// Set env configurations.
	setFromEnv(&cfg)

	return &cfg
}

// Parsing config file.
func parseConfigFile(configPath string) {
	log.Debug().Msgf("Parsing config file: %s", configPath)

	// Split path to folder and file.
	path := strings.Split(configPath, "/")

	viper.AddConfigPath(path[0]) // folder
	viper.SetConfigName(path[1]) // file

	// Read config file.
	if err := viper.ReadInConfig(); err != nil {
		// Set default variables.
		viper.SetDefault("http.addr", defaultHTTPPort)
		viper.SetDefault("http.appName", defaultHTTPAppName)

		log.Error().Msgf("error parsing config file: %s", err.Error())
	}
}

// Unmarshal config keys.
func unmarshal(cfg *Config) {
	log.Debug().Msg("Unmarshal config keys...")

	// Unmarshal http keys.
	if err := viper.UnmarshalKey("http", &cfg.HTTP); err != nil {
		log.Error().Msgf("error unmarshal http keys: %s", err.Error())
	}
	// Unmarshal grpc keys.
	if err := viper.UnmarshalKey("grpc", &cfg.GRPC); err != nil {
		log.Error().Msgf("error unmarshal grpc keys: %s", err.Error())
	}
	// Unmarshal auth service keys.
	if err := viper.UnmarshalKey("service.auth", &cfg.Service.Auth); err != nil {
		log.Error().Msgf("error unmarshal auth service keys: %s", err.Error())
	}
}

// Seting environment variables from .env file.
func setFromEnv(cfg *Config) {
	cfg.Auth.SigningKey = os.Getenv("AUTH_SIGNING_KEY")
}
