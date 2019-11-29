FILES = main.go minesweeper.go http.go

.PHONY: all
all: minesweeper

minesweeper: $(FILES)
	go build

run: $(FILES)
	go run $^

.PHONY: test
test:
	go test

.PHONY: clean
clean:
	go clean