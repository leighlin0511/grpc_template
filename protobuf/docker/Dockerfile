# syntax = docker/dockerfile:1.3

ARG PROTOC_VERSION=3.17.3
ARG GO_VERSION=1.17.1
ARG GRPC_GATEWAY_VERSION=2.5.0
ARG GRPC_VERSION=1.39.1
ARG GRPC_WEB_VERSION=1.2.1
ARG PROTOC_GEN_GO_VERSION=1.27.1
ARG PROTOC_GEN_GO_GRPC_VERSION=1.1.0
ARG PROTOC_GEN_VALIDATE_VERSION=0.6.1

FROM golang:1.17.1-bullseye as builder
WORKDIR /src
RUN apt-get update -y && \
    apt-get dist-upgrade -y && \
    apt-get -y autoremove && \
    apt-get clean 
RUN apt-get install -y \
    git \
    build-essential \
    autoconf \
    libtool \
    curl \
    unzip

ARG PROTOC_VERSION
ENV PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip 
RUN curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/$PROTOC_ZIP \
    && unzip -o $PROTOC_ZIP -d /usr/local bin/protoc \
    && rm -f $PROTOC_ZIP

ARG GRPC_VERSION
RUN go get google.golang.org/grpc@v${GRPC_VERSION}

ARG PROTOC_GEN_GO_VERSION
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v${PROTOC_GEN_GO_VERSION} \
    && export PATH="$PATH:$(go env GOPATH)/bin"

ARG PROTOC_GEN_GO_GRPC_VERSION
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v${PROTOC_GEN_GO_GRPC_VERSION} \
    && export PATH="$PATH:$(go env GOPATH)/bin"

ARG GRPC_GATEWAY_VERSION
RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v${GRPC_GATEWAY_VERSION} \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v${GRPC_GATEWAY_VERSION} \
    && export PATH="$PATH:$(go env GOPATH)/bin"

COPY generate-go.sh generate-go.sh
COPY proto proto

RUN --mount=type=cache,target=/root/.cache/go-build \
    mkdir -p generated && \
    mkdir -p generated/pkg && \
    bash ./generate-go.sh

FROM scratch as generating

LABEL maintainer="Leigh Lin &lt;leighlin0511@hotmail.com&gt;"

COPY --from=builder src/generated/ /
