# op-be-graphql
The GraphQL Services for the Backend Layer of OctopOSPrime

[![Build - Test](https://github.com/octoposprime/op-be-graphql/actions/workflows/ci.yml/badge.svg)](https://github.com/octoposprime/op-be-graphql/actions/workflows/ci.yml)
[![Docker Image Publish](https://github.com/octoposprime/op-be-graphql/actions/workflows/cd.yml/badge.svg)](https://github.com/octoposprime/op-be-graphql/actions/workflows/cd.yml)

## Pre-Requirements
Graphql
```
export PATH="$PATH:$(go env GOPATH)/bin"

go get github.com/99designs/gqlgen

make gql-generate
```

## Development Environment
You have to see ![github.com/octoposprime/op-be-docs](https://github.com/octoposprime/op-be-docs) before development.

#### .env
```
POSTGRES_USERNAME=op
POSTGRES_PASSWORD=op
JWT_SECRET_KEY=op
REDIS_PASSWORD=op
```

#### Local Run
```
make local-run
```

#### Docker Run
```
TEST=true POSTGRES_USERNAME=op POSTGRES_PASSWORD=op JWT_SECRET_KEY=op REDIS_PASSWORD=op make docker-build
make docker-run 


```