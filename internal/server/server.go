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

package server

import (
	"github.com/durudex/durudex-gateway/internal/config"
	"github.com/durudex/durudex-gateway/internal/delivery/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// Server structure.
type Server struct {
	httpApp     *fiber.App
	httpHandler *http.Handler
}

// Creating a new application http server.
func NewServer(cfg *config.Config, httpHandler *http.Handler) *Server {
	return &Server{
		httpApp:     fiber.New(fiber.Config{AppName: cfg.Server.Name}),
		httpHandler: httpHandler,
	}
}

// Running application server.
func (s *Server) Run(addr string) {
	log.Debug().Msg("Running server...")

	// Initialize http routes.
	s.httpHandler.InitRoutes(s.httpApp)

	if err := s.httpApp.Listen(addr); err != nil {
		log.Fatal().Err(err).Msg("error running http application")
	}
}

// Stoping application server.
func (s *Server) Stop() {
	log.Info().Msg("Stoping server...")

	if err := s.httpApp.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("error stoping http application")
	}
}
