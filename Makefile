export MIGRATE_PATH := /Users/hoai/personal/wager_service/tools/migration/mysql/migrations
export MYSQL_URI := root:sieunhanGAobay23vong@tcp(127.0.0.1:3306)
export MYSQL_URI_PARAMS := maxAllowedPacket=0&multiStatements=true&parseTime=true
export MYSQL_TEST_DB := wager_test
export MYSQL_DB := wager

.PHONY: test

test:
	go test ./...

run:
	go run ./tools/migration/main.go