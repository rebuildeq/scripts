SHELL := /bin/bash

# runs build-all
run: build-all
# compilies the project to bin/
build:
	mkdir -p bin
	go build -o bin/rebuildeq main.go
# runs all build commands
build-all: build copy-data
	cd bin && ./rebuildeq build all
	cp bin/aa.md ../web/content/aa.md
	cp bin/dbstr_us.txt ../launcheq/rof
	cp bin/spells_us.txt ../launcheq/rof
# builds aa
build-aa: build copy-data
	cd bin && ./rebuildeq build aa
	cp bin/aa.md ../web/content/aa.md
	cp bin/dbstr_us.txt ../launcheq/rof
	cp bin/spells_us.txt ../launcheq/rof
# builds spell
build-spell: build copy-data
	cd bin && ./rebuildeq build spell
	cp bin/spells_us.txt ../launcheq/rof
	cp bin/dbstr_us.txt ../launcheq/rof
# builds rule
build-rule: build copy-data
	cd bin && ./rebuildeq build rule
# builds task
build-task: build copy-data
	cd bin && ./rebuildeq build task
# builds charcreate
build-charcreate: build copy-data
	cd bin && ./rebuildeq build charcreate
build-item: build copy-data
	cd bin && ./rebuildeq build item
# copies data from data/ to bin/
copy-data:
	cp data/* bin/
	cp .env bin/
import-aa: build copy-data
	cd bin && ./rebuildeq import aa
import-rule: build copy-data
	cd bin && ./rebuildeq import rule
import-spell: build copy-data
	cd bin && ./rebuildeq import spell 23598
import-task: build copy-data
	cd bin && ./rebuildeq import task
import-charcreate: build copy-data
	cd bin && ./rebuildeq import charcreate
import-npc: build copy-data
	cd bin && ./rebuildeq import npc
import-item: build copy-data
	cd bin && ./rebuildeq import item 1001
show-tables: copy-data
	source .env && cd bin && docker run --rm \
	-v ${PWD}:/src \
    imega/mysql-client \
    mysql --host=$$DB_HOST --user=$$DB_USER --password=$$DB_PASS --database=$$DB_NAME \
	--execute='show tables;'
# inject tries to insert the sql files from bin/ to the db
inject: copy-data
	docker run -it --rm \
	-v ${PWD}/bin:/src \
    imega/mysql-client \
	/bin/sh -c 'source /src/.env && mysql --host=$$DB_HOST --user=$$DB_USER --password=$$DB_PASS $$DB_NAME < /src/aa.sql'

