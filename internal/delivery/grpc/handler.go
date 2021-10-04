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

package grpc

import (
	"github.com/Durudex/durudex-gateway/internal/config"
	pb "github.com/Durudex/durudex-gateway/internal/delivery/grpc/protobuf"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Handler struct {
	Auth pb.AuthServiceClient
}

// Creating a new grpc handler.
func NewGRPCHandler(cfg *config.Config) *Handler {
	authServiceConn := ConnectToService(cfg.Service.Auth.Addr, cfg.Service.Auth.CertPath)

	return &Handler{
		Auth: pb.NewAuthServiceClient(authServiceConn),
	}
}

// Connecting to service.
func ConnectToService(address, certPath string) *grpc.ClientConn {
	log.Debug().Msgf("Connecting to %s service", address)

	// Connecting to service.
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(NewGRPCCreds(certPath)))
	if err != nil {
		log.Error().Msgf("error connecting to service: %s", err.Error())
	}

	return conn
}

// New server TLS from file.
func NewGRPCCreds(certPath string) credentials.TransportCredentials {
	creds, err := credentials.NewClientTLSFromFile(certPath, "")
	if err != nil {
		log.Fatal().Msgf("error creating a new server TLS from file: %s", err.Error())
	}

	return creds
}
