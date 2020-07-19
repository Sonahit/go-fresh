run:
	go run main.go

compile:
	mkdir -p build
	go build -o build/main main.go
