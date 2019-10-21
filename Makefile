default: help

DUMPFILE := artifacts/fcc-license-view-data-csv-format.zip
BOLTDB := artifacts/fcc.db
FCC2BOLT := bin/fcc2bolt
FCC := bin/fccdb
VARDIR := /var/local/fccdb

.PHONY: fcc
fcc: $(FCC) $(BOLTDB)

.PHONY: help
help:
	@echo
	@echo "ingest: ingest FCC database into local database format"
	@echo "fcc: build binaries and include ingest"
	@echo "install: copy db file to $(VARDIR) and fcc binary to /usr/local/bin"
	@echo "download: download FCC dataset"
	@echo

.PHONY: ingest
ingest: $(BOLTDB)

.PHONY: clean
clean:
	rm -rf artifacts/*
	rm -rf bin/*

$(VARDIR):
	mkdir -p $@

.PHONY: download
download: $(DUMPFILE)

$(DUMPFILE):
	cd artifacts && wget http://data.fcc.gov/download/license-view/fcc-license-view-data-csv-format.zip

$(BOLTDB): $(DUMPFILE) | $(FCC2BOLT)
	$(FCC2BOLT) -dump $(DUMPFILE) -db $(BOLTDB)

$(FCC2BOLT):
	go build -o $@ cmd/fcc2bolt/main.go

.PHONY: install
install: fcc $(VARDIR)
	cp $(FCC) /usr/local/bin/fcc
	cp $(BOLTDB) $(VARDIR)
 
