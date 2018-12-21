all: run

build:
	go build

run: build
	./part2 < ../input.txt

test: build
	./part2 < ../testInput.txt

watch:
	ag -l . ../ | entr make run

watch-test:
	ag -l . ../ | entr make test
