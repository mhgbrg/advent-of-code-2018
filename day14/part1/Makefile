all: run

build:
	go build

run: build
	./part1 < ../input.txt

test: build
	./part1 < ../testInput.txt

watch:
	ag -l . ../ | entr make run

watch-test:
	ag -l . ../ | entr make test
