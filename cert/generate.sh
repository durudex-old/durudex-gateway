#!/bin/sh

echo "Write password: "
read PASSWORD

echo "Generate CA key..."
openssl genrsa -passout pass:$PASSWORD -des3 -out cert/ca.key 4096

echo "Generate CA certificate..."
openssl req -passin pass:$PASSWORD -new -x509 -days 365 -key cert/ca.key \
  -out cert/ca.crt -subj "/C=UK/ST=KY/L=Durudex/O=Durudex/OU=Durudex/CN=cert/"

echo "Generate authservice key..."
openssl genrsa --passout pass:$PASSWORD -des3 -out cert/authservice.key 4096

echo "Generate authservice signing request..."
openssl req -passin pass:$PASSWORD -new -key cert/authservice.key -out cert/authservice.csr \
  -subj "/C=UK/ST=KY/L=Durudex/O=Durudex/OU=Durudex/CN=%COMPUTERNAME%"

echo "Self-sign authservice certificate..."
openssl x509 -req -passin pass:$PASSWORD -days 365 -in cert/authservice.csr \
  -CA cert/ca.crt -CAkey cert/ca.key -set_serial 01 -out cert/authservice.crt

echo "Remove passphrase from authservice key..."
openssl rsa -passin pass:$PASSWORD -in cert/authservice.key -out cert/authservice.key

echo "Generate client key..."
openssl genrsa -passout pass:$PASSWORD -des3 -out cert/client.key 4096

echo "Generate client signing request..."
openssl req -passin pass:$PASSWORD -new -key cert/client.key -out cert/client.csr \
  -subj "/C=UK/ST=KY/L=Durudex/O=Durudex/OU=Durudex/CN=%CLIENT-COMPUTERNAME%"

echo "Self-sign client certificate..."
openssl x509 -passin pass:$PASSWORD -req -days 365 -in cert/client.csr -CA cert/ca.crt \
  -CAkey cert/ca.key -set_serial 01 -out cert/client.crt

echo "Remove passphrase from client key..."
openssl rsa -passin pass:$PASSWORD -in cert/client.key -out cert/client.key
