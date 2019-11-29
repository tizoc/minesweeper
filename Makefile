FILES = main.go minesweeper.go http.go

.PHONY: all
all: minesweeper

minesweeper: $(FILES)
	go build

run: $(FILES)
	env PORT=8080 go run $^

.PHONY: test
test:
	go test

.PHONY: clean
clean:
	go clean