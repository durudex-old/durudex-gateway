/*
 * Copyright © 2022 Durudex
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

package domain

// Authorization tokens.
type Tokens struct {
	// JWT access token.
	Access string `json:"access"`
	// Refresh token.
	Refresh string `json:"refresh"`
}

// User Sign In input.
type SignInInput struct {
	// Account username.
	Username string `json:"username"`
	// User password
	Password string `json:"password"`
	// User ip address.
	Ip string
}

// User Sign Up input.
type SignUpInput struct {
	// Account username.
	Username string `json:"username"`
	// User email.
	Email string `json:"email"`
	// User password.
	Password string `json:"password"`
	// User verification code.
	Code uint64 `json:"code"`
	// User ip address.
	Ip string
}

// Refresh tokens input.
type RefreshTokenInput struct {
	// Refresh token.
	Token string `json:"token"`
	// User ip address.
	Ip string
}
