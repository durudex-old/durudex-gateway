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
	"fmt"

	"github.com/golang-jwt/jwt"
)

type JWTConfig struct{ SigningKey string }

type JWTManager struct {
	cfg JWTConfig
}

// Creating a new JWT manager.
func NewJWTManager(cfg JWTConfig) *JWTManager {
	return &JWTManager{cfg: cfg}
}

// Parsing jwt access token.
func (m *JWTManager) Parse(accessToken string) (string, error) {
	// Parsing and validation token.
	token, err := jwt.Parse(accessToken, m.keyFunc)
	if err != nil {
		return "", err
	}

	// Get user claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("error get user claims from token")
	}

	return claims["sub"].(string), nil
}

func (m *JWTManager) keyFunc(t *jwt.Token) (i interface{}, err error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
	}

	return []byte(m.cfg.SigningKey), nil
}
