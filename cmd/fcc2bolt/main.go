package main

import (
	"archive/zip"
	"encoding/csv"
	"flag"
	"io"
	"log"
	"strconv"
	"time"

	"github.com/alrs/fcc"

	"github.com/davecgh/go-spew/spew"
	bolt "go.etcd.io/bbolt"
)

const csvFilename = "fcc_lic_vw.csv"

func main() {
	var dumpPath, boltPath string
	newest := time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)

	flag.StringVar(&dumpPath, "dump", "", "path to FCC database dump zipfile")
	flag.StringVar(&boltPath, "db", "./fcc.db", "path to local database destination")
	flag.Parse()

	if dumpPath == "" {
		log.Fatal("location of dump path required")
	}

	db, err := bolt.Open(boltPath, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	tx, err := db.Begin(true)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	lb, err := tx.CreateBucket([]byte("licenses"))
	if err != nil {
		log.Fatal(err)
	}

	zf, err := zip.OpenReader(dumpPath)
	if err != nil {
		log.Fatalf("error opening FCC zip:%s", err)
	}
	defer zf.Close()

	var wCount, cCount, eCount uint
	for _, file := range zf.File {
		if file.FileHeader.Name != csvFilename {
			log.Fatalf("found file %s, archive should only contain %s", file.FileHeader.Name,
				csvFilename)
		}

		f, err := file.Open()
		if err != nil {
			log.Fatalf("error opening file inside zipfile: %s", err)
		}
		defer f.Close()

		reader := csv.NewReader(f)
		_, _ = reader.Read() // ignore header

		for i := 0; ; i++ {
			if (i % 1000000) == 0 {
				log.Printf("%d read / %d wrote / %d cncl / %d expr / newest:%v",
					i, wCount, cCount, eCount, newest)
			}
			line, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("error parsing %s: %s", line, err)
			}
			dl, err := fcc.ParseLicense(line)
			if err != nil {
				log.Fatalf("failed to parse %s: %s", spew.Sdump(dl), err)
			}

			if dl.CancellationDate != nil &&
				dl.CancellationDate.After(newest) &&
				!(dl.CancellationDate.After(time.Now())) {
				newest = *dl.CancellationDate
			}

			if dl.CancellationDate != nil {
				cCount++ // it's canceled
				continue
			}
			if dl.ExpiredDate == nil {
				eCount++
				continue
			}

			if time.Now().Before(*dl.ExpiredDate) &&
				(dl.RadioServiceCode == "HA" || dl.RadioServiceCode == "HV") {
				minimal := dl.Minimal()
				lb.Put([]byte(dl.Callsign), minimal.DiskFormat())
				wCount++
			}

		}
	}

	mb, err := tx.CreateBucket([]byte("metadata"))
	if err != nil {
		log.Fatal(err)
	}

	mb.Put([]byte("format"), []byte("0"))
	mb.Put([]byte("ingestDate"), []byte(time.Now().String()))
	mb.Put([]byte("latestDate"), []byte(newest.String()))
	mb.Put([]byte("recordCount"), []byte(strconv.Itoa(int(wCount))))

	err = tx.Commit()
	if err != nil {
		log.Fatalf("error on commit: %s", err)
	}

}
