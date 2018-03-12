serve: serve-dev

serve-dev:
	HTTP_ADDRESS=127.0.0.1:1234 APP_ENV=dev JWT_PUBLIC_KEY=public.pem go run main.go

serve-test:
	HTTP_ADDRESS=127.0.0.1:1234 APP_ENV=test go run main.go

lint:
	gometalinter ./...

test:
	go test -v ./...