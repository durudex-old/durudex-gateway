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

package grpc

import (
	"github.com/durudex/durudex-gateway/internal/config"
	"github.com/durudex/durudex-gateway/internal/delivery/grpc/pb"
	"github.com/durudex/durudex-gateway/pkg/tls"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	CACertFile     = "certs/rootCA.pem"
	clientCertFile = "certs/client-cert.pem"
	clientKeyFile  = "certs/client-key.pem"
)

// GRPC handler structure.
type Handler struct {
	Auth pb.AuthServiceClient
	Code pb.CodeServiceClient
	User pb.UserServiceClient
}

// Creating a new grpc handler.
func NewGRPCHandler(cfg *config.Config) *Handler {
	return &Handler{
		Auth: pb.NewAuthServiceClient(connectToService(cfg.Service.Auth)),
		Code: pb.NewCodeServiceClient(connectToService(cfg.Service.Code)),
		User: pb.NewUserServiceClient(connectToService(cfg.Service.User)),
	}
}

// Connecting to gRPC service.
func connectToService(cfg config.Service) *grpc.ClientConn {
	log.Debug().Msgf("Connecting to %s service", cfg.Addr)

	var transportOption []grpc.DialOption

	// Check is server TLS.
	if cfg.TLS {
		// Loading server TLS credentials.
		tlsCredentials, err := tls.LoadTLSCredentials(CACertFile, clientCertFile, clientKeyFile)
		if err != nil {
			log.Fatal().Msgf("error load tls credentials: %s", err.Error())
		}

		// Append TLS credentials.
		transportOption = append(transportOption, grpc.WithTransportCredentials(tlsCredentials))
	} else {
		transportOption = append(transportOption, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Connecting to service.
	conn, err := grpc.Dial(cfg.Addr, transportOption...)
	if err != nil {
		log.Error().Msgf("error connecting to service: %s", err.Error())
	}

	return conn
}
