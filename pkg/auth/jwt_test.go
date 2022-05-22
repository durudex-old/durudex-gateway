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

package auth_test

import (
	"reflect"
	"testing"

	"github.com/durudex/durudex-gateway/pkg/auth"
)

// Testing parsing jwt access token.
func Test_Parse(t *testing.T) {
	// Testing args.
	type args struct{ accessToken, signingKey string }

	// Tests structures.
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "OK",
			args: args{
				accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODQzMzI3NjMsInN1YiI6IjEifQ.xEKQVpR4-IGc13wz43LN0TeDfXhBbX57Qe_DVloyJvM",
				signingKey:  "super-key",
			},
			want: "1",
		},
		{
			name:    "Inavalid Access Token",
			args:    args{accessToken: "", signingKey: "super-key"},
			wantErr: true,
		},
	}

	// Conducting tests in various structures.
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Parsing jwt access token
			got, err := auth.Parse(tt.args.accessToken, tt.args.signingKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("error parsing access token: %s", err.Error())
			}

			// Check for similarity of a claims.
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("error claims are not similar: %s", err.Error())
			}
		})
	}
}
