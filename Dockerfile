FROM golang:latest

ARG INSTALLED_LIB="postgresql-client python3 python3-setuptools python3-venv python3-pip"

WORKDIR /app

RUN set -ex \
    && apt-get update \
    && apt-get install --no-install-recommends -y $INSTALLED_LIB \
    && apt-get clean \
    && apt-get autoclean \
    && rm -rf /var/lib/apt/lists/*

RUN set -ex \
    && wget https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz -O /tmp/migrate.linux-amd64.tar.gz \
    && tar -C /tmp/ -xvzf /tmp/migrate.linux-amd64.tar.gz \
    && rm /tmp/migrate.linux-amd64.tar.gz \
    && cp /tmp/migrate.linux-amd64 /usr/local/bin/migrate

RUN pip3 install pre-commit
RUN set -ex \
    && GO111MODULE=on \
    && go get github.com/go-delve/delve/cmd/dlv \
    && go get github.com/githubnemo/CompileDaemon \
    && go get golang.org/x/tools/cmd/goimports \
    && go get github.com/sqs/goreturns \
    && go get golang.org/x/lint/golint \
    && go get github.com/go-critic/go-critic/cmd/gocritic
RUN GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.27.0

COPY go.mod .
COPY go.sum .
RUN go mod download
