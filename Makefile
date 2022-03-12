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

.PHONY: build
build:
	go mod download && CGO_ENABLE=0 GOOS=linux go build -o ./.bin/app ./cmd/gateway/main.go

.PHONY: run
run: build
	docker-compose up --remove-orphans app

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test: lint
	go test -v ./...

.PHONY: gqlgen
gqlgen:
	go get -d github.com/99designs/gqlgen
	go run github.com/99designs/gqlgen generate --config ./gqlgen.yml
	go mod tidy

.PHONY: protoc
protoc:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		internal/delivery/grpc/pb/*.proto

.PHONY: protoc-types
protoc-types:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		internal/delivery/grpc/pb/types/*.proto

.PHONY: cert
cert:
	scripts/generate-cert.sh

.DEFAULT_GOAL := run
