/*
 * Copyright © 2021-2022 Durudex
 *
 * This file is part of Durudex: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * Durudex is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with Durudex. If not, see <https://www.gnu.org/licenses/>.
 */

package config_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/durudex/durudex-gateway/internal/config"
)

// Testing initialize config.
func TestInit(t *testing.T) {
	// Environment configurations.
	type env struct{ configPath, jwtSigningKey string }

	// Testing args.
	type args struct{ env env }

	// Set environment configurations.
	setEnv := func(env env) {
		os.Setenv("CONFIG_PATH", env.configPath)
		os.Setenv("JWT_SIGNING_KEY", env.jwtSigningKey)
	}

	// Default service tls config.
	serviceTLS := config.TLSConfig{
		Enable: true,
		CACert: "./certs/rootCA.pem",
		Cert:   "./certs/client-cert.pem",
		Key:    "./certs/client-key.pem",
	}

	// Testing structures.
	tests := []struct {
		name    string
		args    args
		want    *config.Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{env: env{configPath: "fixtures/main", jwtSigningKey: "super-key"}},
			want: &config.Config{
				HTTP: config.HTTPConfig{
					Host: "api.durudex.local",
					Port: "8000",
					Name: "Durudex API Gateway",
				},
				GraphQL: config.GraphQLConfig{ComplexityLimit: 500},
				Auth:    config.AuthConfig{JWT: config.JWTConfig{SigningKey: "super-key"}},
				Service: config.ServiceConfig{
					User: config.Service{
						Addr: "user.service.durudex.local:8004",
						TLS:  serviceTLS,
					},
					Post: config.Service{
						Addr: "post.service.durudex.local:8005",
						TLS:  serviceTLS,
					},
				},
			},
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment configurations.
			setEnv(tt.args.env)

			// Initialize config.
			got, err := config.Init()
			if (err != nil) != tt.wantErr {
				t.Errorf("error initialize config: %s", err.Error())
			}

			// Check for similarity of a config.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error config are not similar: %s", err.Error())
			}
		})
	}
}
