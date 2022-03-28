<div align="center">
    <a href="https://discord.gg/4qcXbeVehZ">
        <img alt="Discord" src="https://img.shields.io/discord/882288646517035028?label=%F0%9F%92%AC%20discord">
    </a>
    <a href="https://github.com/durudex/durudex-gateway/blob/main/COPYING">
        <img alt="License" src="https://img.shields.io/github/license/durudex/durudex-gateway?label=%F0%9F%93%95%20license">
    </a>
    <a href="https://github.com/durudex/durudex-gateway/stargazers">
        <img alt="GitHub Stars" src="https://img.shields.io/github/stars/durudex/durudex-gateway?label=%E2%AD%90%20stars&logo=sdf">
    </a>
    <a href="https://github.com/durudex/durudex-gateway/network">
        <img alt="GitHub Forks" src="https://img.shields.io/github/forks/durudex/durudex-gateway?label=%F0%9F%93%81%20forks">
    </a>
</div>

<h1 align="center">⚡️ Durudex Gateway</h1>

<p align="center">
GraphQL API Gateway that integrates many services.
</p>

### 💡 Prerequisites
+ [Go 1.17](https://golang.org/)
+ [Docker](https://www.docker.com)
+ [grpc](https://grpc.io/docs/languages/go/quickstart/)
+ [mkcert](https://github.com/FiloSottile/mkcert)
+ [golangci-lint](https://golangci-lint.run/usage/install/)

## ⚙️ Build & Run
1) Add local domains:
```sh
api.durudex.local
```
2) Create an `.env` file in the root directory and add the following values from `.env.example`:
```env
# Config variables:
CONFIG_PATH=configs/main

# Auth variables:
JWT_SIGNING_KEY=
```
3) Generate certificates, information can be found at [certs/README.md](cert/README.md).
4) Run services:
+ [durudex-auth-service](https://github.com/durudex/durudex-auth-service)
+ [durudex-code-service](https://github.com/durudex/durudex-code-service)

Use `make run` to run and `make build` to build project.

## 🛠 Lint & Tests
Use `make lint` to run the lint, and use `make test` for tests.

## 👍 Contribute
If you want to say thank you and/or support the active development of [Durudex](https://github.com/durudex):
1) Add a [GitHub Star](https://github.com/Durudex/durudex-gateway/stargazers) to the project.
2) Join the [Discord Server](https://discord.gg/4qcXbeVehZ).

## ⚠️ License
Copyright © 2021-2022 [Durudex](https://github.com/Durudex). Released under the [GNU AGPL v3](https://www.gnu.org/licenses/agpl-3.0.html) license.

#### Third-party library licenses
+ [fiber](https://github.com/gofiber/fiber/blob/master/LICENSE)
+ [zerolog](https://github.com/rs/zerolog/blob/master/LICENSE)
+ [viper](https://github.com/spf13/viper/blob/master/LICENSE)
+ [gqlgen](https://github.com/99designs/gqlgen/blob/master/LICENSE)
+ [gqlparser](https://github.com/vektah/gqlparser/blob/master/LICENSE)
+ [adaptor](https://github.com/gofiber/adaptor/blob/master/LICENSE)
+ [grpc-go](https://github.com/grpc/grpc-go/blob/master/LICENSE)
+ [protobuf](https://github.com/protocolbuffers/protobuf/blob/master/LICENSE)
+ [jwt](https://github.com/golang-jwt/jwt/blob/main/LICENSE)
