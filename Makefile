buildAll: clean test build buildWindows

build:
	go build

buildWindows:
	env GOOS=windows GOARCH=amd64 go build -o go-soft4fx.exe

test:
	go test  ./...

clean:
	go clean
	rm -f go-soft4fx go-soft4fx.exe