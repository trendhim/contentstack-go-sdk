#!/usr/bin/env bash

set -e

cd "$(git rev-parse --show-toplevel)"

go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen

echo ""
echo "=> Generating delivery files"
echo "   Types:  delivery/types.go"
oapi-codegen \
  --generate types \
  --package delivery \
  delivery.json \
> delivery/types.go

echo "   Client: delivery/client.go"
oapi-codegen \
  --generate client \
  --package delivery \
  delivery.json \
> delivery/client.go

echo ""
echo "=> Generating management files"
echo "   Types:  management/types.go"
oapi-codegen \
  --generate types \
  --package management \
  management.json \
> management/types.go

echo "   Client: management/client.go"
oapi-codegen \
  --generate client \
  --package management \
  management.json \
> management/client.go

echo ""
