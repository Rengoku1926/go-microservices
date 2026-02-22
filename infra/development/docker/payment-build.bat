
set CGO_ENABLED=0
set GOOS=linux
set GOARCH=arm64
go build -o build/payment-service ./services/payment-service/cmd/main.go