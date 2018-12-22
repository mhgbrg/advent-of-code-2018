all: run

build:
	go build

run: build
	./part1 < ../input.txt

test: build
	./part1 < ../testInput.txt

test1: build
	./part1 < ../testInput1.txt

test2: build
	./part1 < ../testInput2.txt

watch:
	ag -l . ../ | entr make run

watch-test:
	ag -l . ../ | entr make test

watch-test1:
	ag -l . ../ | entr make test1

watch-test2:
	ag -l . ../ | entr make test2
