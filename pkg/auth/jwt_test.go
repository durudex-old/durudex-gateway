/*
	Copyright © 2022 Durudex

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

package auth

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
)

// Testing a parsing jwt access token.
func TestParse(t *testing.T) {
	// Creating a new manager.
	jwtManager := NewJWTManager(JWTConfig{SigningKey: "super-key"})

	// Create a new access token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
		Subject:   "1234567890",
	})
	accessToken, err := token.SignedString([]byte(jwtManager.cfg.SigningKey))
	if err != nil {
		t.Errorf("error creating a new access token: %s", err.Error())
	}

	// Parsing jwt access token.
	val, err := jwtManager.Parse(accessToken)
	if err != nil {
		t.Errorf("error parsing access token: %s", err.Error())
	}

	t.Logf("Access token value: %s", val)
}
