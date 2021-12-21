#!/bin/sh

# Copyright © 2021 Durudex

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

cd cert/

export CAROOT="./"

if [ -f './rootCA.pem' ]; then
  echo "Delete old CA and certificates..."
  mkcert -uninstall
  rm *.pem
fi

echo "Install new CA..."
mkcert -install

echo "Create Auth Service certificate...."
mkcert \
  -cert-file authservice-cert.pem \
  -key-file authservice-key.pem \
  authservice.durudex.local

echo "Create Notif Service certificate...."
mkcert \
  -cert-file notifservice-cert.pem \
  -key-file notifservice-key.pem \
  notifservice.durudex.local

echo "Create User Service certificate...."
mkcert \
  -cert-file userservice-cert.pem \
  -key-file userservice-key.pem \
  userservice.durudex.local

echo "Create Client certificate...."
mkcert -client \
  -cert-file client-cert.pem \
  -key-file client-key.pem \
  localhost 127.0.0.1
