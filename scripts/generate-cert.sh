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
rm *.crt *.csr *.key

echo "Wrire CA password:"
read password

echo "Generate CA key..."
openssl genrsa -des3 \
  -passout pass:"$password" \
  -out ca.key 4096

echo "Generate CA certificate..."
openssl req -x509 -new -nodes -sha256 \
  -passin pass:"$password" \
  -days 365 \
  -config ca.conf \
  -key ca.key \
  -out ca.crt

echo "Wrire Auth Service password:"
read password

echo "Generate auth service key..."
openssl genrsa -des3 \
  -passout pass:"$password" \
  -out authservice.key 4096

echo "Generate authservice signing request..."
openssl req -new \
  -passin pass:"$password" \
  -config authservice.conf \
  -key authservice.key \
  -out authservice.csr

echo "Self-sign authservice certificate..."
openssl x509 -req \
  -passin pass:"$password" \
  -days 365 \
  -in authservice.csr \
  -extensions 'req_ext' \
  -CA ca.crt \
  -CAkey ca.key \
  -set_serial 01 \
  -out authservice.crt

echo "Remove passphrase from authservice key..."
openssl rsa \
  -passin pass:"$password" \
  -in authservice.key \
  -out authservice.key

echo "Wrire Client password:"
read password

echo "Generate client key..."
openssl genrsa -des3 \
  -passout pass:"$password" \
  -out client.key 4096

echo "Generate client signing request..."
openssl req -new \
  -passin pass:"$password" \
  -config client.conf \
  -key client.key \
  -out client.csr

echo "Self-sign client certificate..."
openssl x509 -req \
  -passin pass:"$password" \
  -days 365 \
  -in client.csr \
  -extensions 'req_ext' \
  -CA ca.crt \
  -CAkey ca.key \
  -set_serial 01 \
  -out client.crt

echo "Remove passphrase from client key..."
openssl rsa \
  -passin pass:"$password" \
  -in client.key \
  -out client.key
