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
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"

	"github.com/Durudex/durudex-gateway/internal/config"
	pb "github.com/Durudex/durudex-gateway/internal/delivery/grpc/protobuf"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	serverCACertFile = "cert/ca-cert.pem"
	clientCertFile   = "cert/client-cert.pem"
	clientKeyFile    = "cert/client-key.pem"
)

type Handler struct {
	Auth pb.AuthServiceClient
}

// Creating a new grpc handler.
func NewGRPCHandler(cfg *config.Config) *Handler {
	transportOption := grpc.WithInsecure()

	// If TLS is true.
	if cfg.GRPC.TLS {
		tlsCredentials, err := LoadTLSCredentials()
		if err != nil {
			log.Fatal().Msgf("error load tls credentials: %s", err.Error())
		}

		transportOption = grpc.WithTransportCredentials(tlsCredentials)
	}

	authServiceConn := ConnectToService(cfg.Service.Auth.Addr, transportOption)

	return &Handler{
		Auth: pb.NewAuthServiceClient(authServiceConn),
	}
}

// Connecting to service.
func ConnectToService(address string, transportOption grpc.DialOption) *grpc.ClientConn {
	log.Debug().Msgf("Connecting to %s service", address)

	// Connecting to service.
	conn, err := grpc.Dial(address, transportOption)
	if err != nil {
		log.Error().Msgf("error connecting to service: %s", err.Error())
	}

	return conn
}

// Loading TLS credentials.
func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate od the CA who signed server's certificate.
	pemServerCA, err := ioutil.ReadFile(serverCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, errors.New("error to add server CA's certificate")
	}

	// Load client's certificate and private key.
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and returning it.
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
