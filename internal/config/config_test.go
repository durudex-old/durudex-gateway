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

package config

import (
	"os"
	"reflect"
	"testing"
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

	// Testing structures.
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "OK",
			args: args{env: env{configPath: "fixtures/main", jwtSigningKey: "super-key"}},
			want: &Config{
				Server: ServerConfig{
					Host: defaultServerHost,
					Port: defaultServerPort,
					Name: defaultServerName,
				},
				Auth: AuthConfig{JWT: JWTConfig{SigningKey: "super-key"}},
				Service: ServiceConfig{
					Auth: Service{
						Addr: defaultServiceAuthAddr,
						TLS: TLSConfig{
							Enable: true,
							CACert: "./certs/rootCA.pem",
							Cert:   "./certs/client-cert.pem",
							Key:    "./certs/client-key.pem",
						},
					},
					Code: Service{
						Addr: defaultServiceCodeAddr,
						TLS: TLSConfig{
							Enable: true,
							CACert: "./certs/rootCA.pem",
							Cert:   "./certs/client-cert.pem",
							Key:    "./certs/client-key.pem",
						},
					},
					User: Service{
						Addr: defaultServiceUserAddr,
						TLS: TLSConfig{
							Enable: true,
							CACert: "./certs/rootCA.pem",
							Cert:   "./certs/client-cert.pem",
							Key:    "./certs/client-key.pem",
						},
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
			got, err := Init()
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
