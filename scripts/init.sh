#!/bin/bash

module_name='gin-api-template'

cd "$(dirname $0)/../"
touch ./build/package/development/secrets.env

sed -i -r "s/gin-api-template/$module_name/g" \
  ./go.mod ./pkg/database/database.go ./pkg/server/server.go ./pkg/server/routes.go \
  ./scripts/alias.sh

docker-compose -f deployments/development/docker-compose.yml -p "$project_name-development" build
