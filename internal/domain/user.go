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

import (
	"time"

	"github.com/segmentio/ksuid"
)

// User type.
type User struct {
	// User id.
	Id ksuid.KSUID `json:"id"`
	// Username.
	Username string `json:"username"`
	// User last visit date.
	LastVisit time.Time `json:"lastVisit"`
	// User verified status.
	Verified bool `json:"verified"`
	// User avatar url.
	AvatarUrl *string `json:"avatarUrl"`
}

func (User) IsNode() {}

// Forgot user password input.
type ForgotPasswordInput struct {
	// User email.
	Email string `json:"email"`
	// New user password.
	Password string `json:"password"`
	// User verification code.
	Code uint64 `json:"code"`
}
