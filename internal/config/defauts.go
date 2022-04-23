/*
 * Copyright © 2022 Durudex

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
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const (
	// Config defaults.
	defaultConfigPath string = "configs/main"

	// Server defaults.
	defaultServerHost string = "api.durudex.local"
	defaultServerPort string = "8000"
	defaultServerName string = "durudex-api-gateway"

	// Auth service defaults.
	defaultServiceAuthAddr string = "auth.service.durudex.local:8001"

	// Code service defaults.
	defaultServiceCodeAddr string = "code.service.durudex.local:8003"

	// User service defaults.
	defaultServiceUserAddr string = "user.service.durudex.local:8004"

	// Postgres defaults.
	defaultServicePostAddr string = "post.service.durudex.local:8005"
)

// Populate defaults config variables.
func populateDefaults() {
	log.Debug().Msg("Populate defaults config variables.")

	// Server defaults.
	viper.SetDefault("server.host", defaultServerHost)
	viper.SetDefault("server.port", defaultServerPort)
	viper.SetDefault("server.name", defaultServerName)

	// Auth service defaults.
	viper.SetDefault("service.auth.addr", defaultServiceAuthAddr)

	// Code service defaults.
	viper.SetDefault("service.code.addr", defaultServiceCodeAddr)

	// User service defaults.
	viper.SetDefault("service.user.addr", defaultServiceUserAddr)

	// Post service defaults.
	viper.SetDefault("service.post.addr", defaultServicePostAddr)
}
