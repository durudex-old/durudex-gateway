/*
	Copyright © 2021-2022 Durudex

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

	"github.com/durudex/durudex-gateway/internal/config"
	"github.com/durudex/durudex-gateway/internal/delivery/graphql"
	"github.com/durudex/durudex-gateway/internal/delivery/grpc"
	"github.com/durudex/durudex-gateway/internal/delivery/http"
	"github.com/durudex/durudex-gateway/internal/server"
	"github.com/durudex/durudex-gateway/internal/service"
	"github.com/durudex/durudex-gateway/pkg/auth"

	"github.com/rs/zerolog/log"
)

// Run durudex gateway application.
func Run() {
	// Initialize config.
	cfg, err := config.Init()
	if err != nil {
		log.Error().Msg(err.Error())
	}

	// Managers
	auth := auth.NewAuthManager(auth.JWTConfig{
		SigningKey: cfg.Auth.JWT.SigningKey,
	})

	// Service, Handlers
	grpcHandler := grpc.NewGRPCHandler(cfg)
	service := service.NewService(grpcHandler)
	graphqlHandler := graphql.NewHandler(service, auth)
	httpHandler := http.NewHTTPHandler(graphqlHandler)

	// Create and run server.
	srv := server.NewServer(cfg, httpHandler)
	addr := cfg.Server.Host + ":" + cfg.Server.Port

	go srv.Run(addr)

	// Quit in application.
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	// Stoping server.
	srv.Stop()

	log.Info().Msg("Durudex Gateway stoping!")
}
