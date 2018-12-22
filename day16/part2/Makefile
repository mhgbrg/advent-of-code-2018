all: run

build:
	go build

run: build
	./part2

watch:
	ag -l . ../ | entr make run
