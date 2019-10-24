default: all

DUMPFILE := artifacts/fcc-license-view-data-csv-format.zip
BOLTDB := artifacts/fcc.db
FCC2BOLT := bin/fcc2bolt
FCCDB := bin/fccdb
DBDIR := /usr/share/fccdb

.PHONY: all
all: $(FCCDB) $(BOLTDB)

.PHONY: help
help:
	@echo
	@echo "all: build binaries and include ingest"
	@echo "ingest: ingest FCC database into local database format"
	@echo "install: copy db file to $(DBDIR) and fcc binary to /usr/local/bin"
	@echo "download: download FCC dataset"
	@echo

.PHONY: ingest
ingest: $(BOLTDB)

.PHONY: clean
clean:
	rm $(DUMPFILE) $(BOLTDB) $(FCC2BOLT) $(FCCDB)

.PHONY: download
download: $(DUMPFILE)

.PHONY: install
install: all $(DBDIR)
	cp $(FCCDB) /usr/local/bin/fccdb
	cp $(BOLTDB) $(DBDIR)

$(DUMPFILE):
	cd artifacts && wget -4 http://data.fcc.gov/download/license-view/fcc-license-view-data-csv-format.zip

$(BOLTDB): $(DUMPFILE) | $(FCC2BOLT)
	$(FCC2BOLT) -dump $(DUMPFILE) -db $(BOLTDB)

$(FCC2BOLT):
	go build -o $@ cmd/fcc2bolt/main.go

$(FCCDB):
	go build -o $@ cmd/fccdb/main.go

$(DBDIR):
	mkdir -p $@
