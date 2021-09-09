#!/usr/bin/env bash

set -e

cd "$(git rev-parse --show-toplevel)"

echo ""
echo "=> Downloading OpenAPI Specs"
echo "   Delivery API:   delivery.json"
curl -s 'https://assets.contentstack.io/v3/assets/bltc5a09bf374882538/blt318e89ee28df7712/600a93fc0839e910126d6026/cda-openapi.json?disposition=download' \
  | sed 's/\\\//\//g' \
  > deliveryx.json

echo "   Management API: management.json"
curl -s 'https://assets.contentstack.io/v3/assets/bltc5a09bf374882538/blt35dc5eaeb2cfc52d/600a93fd4485e50f8f209efe/cma-openapi.json?disposition=download' \
  | sed 's/\\\//\//g' \
  > managementx.json

echo ""
