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

printf "Enter domain name: "
read domain

if [ -z "$domain" ]; then
  echo "Domain name is empty!"
  exit 1
fi

mkcert -install

mkcert \
  -cert-file auth.service.$domain-cert.pem \
  -key-file auth.service.$domain-key.pem \
  auth.service.$domain

mkcert \
  -cert-file email.service.$domain-cert.pem \
  -key-file email.service.$domain-key.pem \
  email.service.$domain

mkcert \
  -cert-file user.service.$domain-cert.pem \
  -key-file user.service.$domain-key.pem \
  user.service.$domain

mkcert \
  -cert-file code.service.$domain-cert.pem \
  -key-file code.service.$domain-key.pem \
  code.service.$domain

mkcert \
  -cert-file post.service.$domain-cert.pem \
  -key-file post.service.$domain-key.pem \
  post.service.$domain

mkcert -client \
  -cert-file client-cert.pem \
  -key-file client-key.pem \
  localhost 127.0.0.1 api.$domain
