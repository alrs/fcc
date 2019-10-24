# fcc

The root `fcc` package provides the data structures necessary for parsing the
amateur-relevant portion of the FCC license database that is provided as a
1GB+ csv file.

`cmd/fcc2bolt` is a console application that reads the zipfile of the FCC dump,
parses it, and inserts active amateur licenses into a boltDB store.

`cmd/fccdb` is a console application for doing name and address lookups of
FCC-issued amateur callsigns.

## Installation
The `fcc` package can be used as-is. To download and ingest the FCC database,
this software requires `wget`, `make`, and the `go` toolchain.

`make help` provides all options in the Makefile.

`make` will build the `cmd` binaries, download the FCC dataset, and perform
the boltdb ingestion.

`sudo make install` will copy `fccdb` to `/usr/local/bin/` and the `fcc.db`
database to `/usr/share/fccdb`.
