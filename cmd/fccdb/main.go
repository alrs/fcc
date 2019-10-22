// fccdb
// Copyright (C) 2019 Lars Lehtonen KJ6CBE

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/alrs/fcc"
	bolt "go.etcd.io/bbolt"
)

var dbPaths = []string{
	"/usr/share/fccdb/fcc.db",
	"./fcc.db",
}

func main() {
	var path string
	var license, metadata bool
	flag.BoolVar(&license, "l", false, "print license")
	flag.BoolVar(&metadata, "m", false, "print metadata")
	flag.Parse()

	callsign := strings.ToUpper(flag.Arg(0))

	if license {
		fmt.Print(
			`
Copyright (C) 2019 Lars Lehtonen KJ6CBE

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

`)
		os.Exit(0)
	}

	if os.Getenv("FCCDB") != "" {
		path = os.Getenv("FCCDB")
	} else {
		var err error
		path, err = pickPath()
		if err != nil {
			log.Fatalf("FCCDB unset, %s", err)
		}
	}

	options := &bolt.Options{
		ReadOnly: true,
	}

	db, err := bolt.Open(path, 0444, options)
	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}
	defer db.Close()

	tx, err := db.Begin(false)
	if err != nil {
		log.Fatalf("error starting transaction: %s", err)
	}
	defer tx.Rollback()

	if metadata {
		printMetadata(path, tx)
		os.Exit(0)
	}

	if callsign == "" {
		fmt.Printf("\nfccdb <CALLSIGN>\nprints the address to stdout\n\n")
		flag.Usage()
		os.Exit(2)
	}

	b := tx.Bucket([]byte("licenses"))
	rawRecord := b.Get([]byte(callsign))
	if rawRecord == nil {
		log.Fatalf("no record found: %s", callsign)
	}
	record := fcc.ReadRecord(rawRecord)
	printRecord(callsign, record)
}

func printMetadata(path string, tx *bolt.Tx) {
	fmt.Printf("using database at path: %s\n\n", path)
	meta := tx.Bucket([]byte("metadata"))
	meta.ForEach(func(k, v []byte) error {
		fmt.Printf("%s:\t%s\n", string(k), string(v))
		return nil
	})
	return
}

func printRecord(call string, r fcc.MinimalLicense) {
	fmt.Printf("\n%s\t%s\n", call, r.Name)
	fmt.Printf("\t%s\n", r.Address)
	fmt.Printf("\t%s, %s %s\n\n", r.City, r.State, r.ZIP)
	return
}

func pickPath() (string, error) {
	for _, path := range dbPaths {
		_, err := os.Stat(path)
		if os.IsNotExist(err) {
			continue
		}
		return path, nil
	}
	return "", fmt.Errorf("could not find fcc.db in %s",
		strings.Join(dbPaths, ","))
}
