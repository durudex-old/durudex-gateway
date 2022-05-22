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

package http

import (
	"github.com/durudex/durudex-gateway/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// HTTP server structure.
type Server struct {
	app *fiber.App
	cfg config.HTTPConfig
}

// Creating a new application http server.
func NewServer(cfg config.HTTPConfig) *Server {
	return &Server{
		app: fiber.New(fiber.Config{AppName: cfg.Name}),
		cfg: cfg,
	}
}

// Running application http server.
func (s *Server) Run() {
	log.Debug().Msg("Running http server...")

	addr := s.cfg.Host + ":" + s.cfg.Port

	// Listen serves HTTP requests from the given addr.
	if err := s.app.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("failed to start http server")
	}
}

// Stopping application http server.
func (s *Server) Stop() {
	log.Info().Msg("Stopping http server...")

	if err := s.app.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("failed to stop http server")
	}
}
