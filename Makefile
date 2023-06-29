# runs build-all
run: build-all
# compilies the project to bin/
build:
	mkdir -p bin
	go build -o bin/rebuildeq main.go
# runs all build commands
build-all:
	cd bin && ./rebuildeq build all
# builds aa
build-aa: build copy-data
	cd bin && ./rebuildeq build aa
# builds spell
build-spell: build copy-data
	cd bin && ./rebuildeq build spell
# builds rule
build-rule: build copy-data
	cd bin && ./rebuildeq build rule
# copies data from data/ to bin/
copy-data:
	cp data/* bin/
import-rule: build copy-data
	cd bin && ./rebuildeq import rule
import-spell: build copy-data
	cd bin && ./rebuildeq import spell