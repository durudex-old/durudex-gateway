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

package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

type Server struct {
	httpApp *fiber.App
}

// Creating a new server.
func NewServer() *Server {
	return &Server{
		httpApp: fiber.New(fiber.Config{}),
	}
}

// Run the server at the specified address.
func (s *Server) Run(addr string) {
	log.Debug().Msg("Running server...")

	if err := s.httpApp.Listen(addr); err != nil {
		log.Fatal().Msgf("error running http application: %s", err.Error())
	}
}

// Server stop.
func (s *Server) Stop() {
	log.Info().Msg("Stoping server...")

	if err := s.httpApp.Shutdown(); err != nil {
		log.Fatal().Msgf("error stoping http application: %s", err.Error())
	}
}
