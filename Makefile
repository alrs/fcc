default: all

DUMPFILE := artifacts/fcc-license-view-data-csv-format.zip
BOLTDB := artifacts/fcc.db
FCC2BOLT := bin/fcc2bolt
FCCD := bin/fccd
FCCDB := bin/fccdb
DBDIR := /usr/share/fccdb

.PHONY: all
all: binaries | ingest

.PHONY: help
help:
	@echo
	@echo "all: build binaries and ingest database (default)"
	@echo "binaries: build fccdb and fcc2bolt"
	@echo "ingest: download FCC database and insert relevant records into boltdb"
	@echo "install: copy db file to $(DBDIR) and fcc binary to /usr/local/bin"
	@echo "download: download FCC dataset"
	@echo

.PHONY: binaries
binaries: $(FCC2BOLT) $(FCCDB) $(FCCD)

.PHONY: ingest
ingest: $(BOLTDB)

.PHONY: clean
clean:
	rm -f $(DUMPFILE) $(BOLTDB) $(FCC2BOLT) $(FCCDB) $(FCCD)

.PHONY: download
download: $(DUMPFILE)

.PHONY: install
install: all $(DBDIR)
	cp $(FCCDB) /usr/local/bin/fccdb
	cp $(FCCD) /usr/local/bin/fccd
	cp $(BOLTDB) $(DBDIR)

$(DUMPFILE):
	cd artifacts && wget -4 http://data.fcc.gov/download/license-view/fcc-license-view-data-csv-format.zip

$(BOLTDB): $(DUMPFILE) | $(FCC2BOLT)
	$(FCC2BOLT) -dump $(DUMPFILE) -db $(BOLTDB)

$(FCC2BOLT):
	go build -o $@ cmd/fcc2bolt/main.go

$(FCCDB):
	go build -o $@ cmd/fccdb/main.go

$(FCCD):
	go build -o $@ cmd/fccd/main.go

$(DBDIR):
	mkdir -p $@
