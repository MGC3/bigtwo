all: build 

test: test_game
	echo "All tests done"

test_game:
	cd app/game && go test -v

build: fmt
	echo "Building backend"
	go build

run: build
	echo "Starting backend server..."
	./backend

fmt:
	echo "Formatting..."
	gofmt -l -s -w .
