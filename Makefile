boilerplate:
	mkdir day${DAY} day${DAY}/part1 day${DAY}/part2
	cat TemplateMakefile | sed 's/{PART}/1/g' > day${DAY}/part1/Makefile
	cat TemplateMakefile | sed 's/{PART}/2/g' > day${DAY}/part2/Makefile
	cp template.go day${DAY}/part1/main.go
	cp template.go day${DAY}/part2/main.go
	curl https://adventofcode.com/2018/day/${DAY}/input --header "Cookie: session=<session-id>" > day${DAY}/input.txt
	touch day${DAY}/testInput.txt
