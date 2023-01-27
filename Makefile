all: go-build \
	docker-build \
	docker-save \
	docker-clean

go-build:
	go build -v -buildvcs=false -ldflags="-X 'main.version=$$(git rev-parse --short HEAD)' -X 'main.build=$$(date --iso-8601=seconds)'" -o ./dist/app-dist ./cmd/server

docker-build:
	docker build -f ./build/Dockerfile.local \
	-t zercle/promtpay-qr-services:latest \
	--pull \
	.

docker-save:
	docker save zercle/promtpay-qr-services | gzip > dist/zercle-promtpay-qr-services.tar.gz

docker-clean:
	docker image prune -f

go-mod:
	go get -v -u && go mod tidy

go-test:
	go test -v ./...

go-run:
	go run ./cmd/server/