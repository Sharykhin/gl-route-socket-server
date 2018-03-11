serve: serve-dev

serve-dev:
	APP_ENV=dev JWT_PUBLIC_KEY=public.pem go run main.go

serve-test:
	APP_ENV=test go run main.go

lint:
	gometalinter ./...

test:
	go test -v ./...