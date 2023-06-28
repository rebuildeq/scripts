
run: build copy-data
	cd bin && ./rebuildeq build all
build:
	mkdir -p bin
	go build -o bin/rebuildeq main.go
all: run
aa: build copy-data
	cd bin && ./rebuildeq build aa
spell: build copy-data
	cd bin && ./rebuildeq build spell
rule: build copy-data
	cd bin && ./rebuildeq build rule
copy-data:
	cp data/* bin/