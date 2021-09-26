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

package gateway

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Durudex/durudex-gateway/internal/config"
	"github.com/Durudex/durudex-gateway/internal/delivery/graphql"
	"github.com/Durudex/durudex-gateway/internal/delivery/http"
	"github.com/Durudex/durudex-gateway/internal/server"
	"github.com/Durudex/durudex-gateway/internal/service"
	"github.com/rs/zerolog/log"
)

// Run durudex gateway application.
func Run(configPath string) {
	// Initialize config.
	cfg := config.Init(configPath)

	// Service, Handlers
	service := service.NewService()
	graphqlHandler := graphql.NewGraphQLHandler(service)
	httpHandler := http.NewHTTPHandler(graphqlHandler)

	// Create and run server.
	srv := server.NewServer(cfg, httpHandler)
	go func() {
		srv.Run(cfg.HTTP.Addr)
	}()

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stoping server.
	srv.Stop()

	log.Info().Msg("Durudex Gateway stoping!")
}
