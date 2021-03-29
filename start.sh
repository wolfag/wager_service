#!/bin/sh
docker-compose -f docker-composer.yml up -d --build

# run test
#MIGRATE_PATH=$(pwd)/tools/migration/mysql/migrations go test ./...
#rc=$?
#if [ $rc -ne 0 ]; then
#  echo "test failed" >&2
#  exit $rc
#fi