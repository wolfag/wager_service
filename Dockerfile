FROM golang:1.16.2 AS builder

WORKDIR /wager_service

ENV GO111MODULE=on
COPY go.mod go.sum ./
RUN go mod download

FROM builder as test
#ENV MIGRATE_PATH="/wager_service/tools/migration/mysql/migrations"
#ENV MYSQL_URI=root:sieunhanGAobay23vong@tcp(wager_mysql:3306)
#ENV MYSQL_URI_PARAMS="maxAllowedPacket=0&multiStatements=true&parseTime=true"
#ENV MYSQL_TEST_DB=wager_test
#ENV MYSQL_DB=wager
COPY . .
RUN go build -o ./bin/migration ./tools/migration/main.go
RUN go build -o ./bin/wager_service ./cmd/main.go

#
CMD ./tools/docker/wait-for-it.sh wager_mysql:3306 && go test ./... && ./bin/migration && ./bin/wager_service

