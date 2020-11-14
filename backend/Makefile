all: test build run

test: test_game
	echo "All tests done"

test_game:
	cd app/game && go test -v

build:
	echo "Building backedn"
	go build

run: build
	echo "Starting backend server..."
	./backend
