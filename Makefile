default: help

DUMPFILE := artifacts/fcc-license-view-data-csv-format.zip
BOLTDB := artifacts/fcc.db

.PHONY: help
help:
	@echo
	@echo "ingest: ingest FCC database into local database format"
	@echo

$(DUMPFILE):
	cd artifacts && wget http://data.fcc.gov/download/license-view/fcc-license-view-data-csv-format.zip

.PHONY: clean
clean:
	rm -rf artifacts/*

$(BOLTDB): $(DUMPFILE)
	go run ./cmd/fcc2bolt/main.go -dump $(DUMPFILE) -db $(BOLTDB)

.PHONY: ingest
ingest: $(BOLTDB)
