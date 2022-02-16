#!/bin/sh

# Copyright © 2021-2022 Durudex

# This file is part of Durudex: you can redistribute it and/or modify
# it under the terms of the GNU Affero General Public License as
# published by the Free Software Foundation, either version 3 of the
# License, or (at your option) any later version.

# Durudex is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
# GNU Affero General Public License for more details.

# You should have received a copy of the GNU Affero General Public License
# along with Durudex. If not, see <https://www.gnu.org/licenses/>.

cd certs/

export CAROOT="./"

if [ -f './rootCA.pem' ]; then
  echo "Delete old CA and certificates..."
  mkcert -uninstall
  rm *.pem
fi

echo "Install new CA..."
mkcert -install

echo "Create Auth Service certificate..."
mkcert \
  -cert-file auth.service-cert.pem \
  -key-file auth.service-key.pem \
  auth.service.durudex.local

echo "Create Email Service certificate..."
mkcert \
  -cert-file email.service-cert.pem \
  -key-file email.service-key.pem \
  email.service.durudex.local

echo "Create User Service certificate..."
mkcert \
  -cert-file user.service-cert.pem \
  -key-file user.service-key.pem \
  user.service.durudex.local

echo "Create Code Service certificate..."
mkcert \
  -cert-file code.service-cert.pem \
  -key-file code.service-key.pem \
  code.service.durudex.local

echo "Create Client certificate..."
mkcert -client \
  -cert-file client-cert.pem \
  -key-file client-key.pem \
  localhost 127.0.0.1 api.durudex.local
