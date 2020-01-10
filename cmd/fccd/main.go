package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	bolt "go.etcd.io/bbolt"

	"github.com/alrs/fcc"
)

func licenseHandler(w http.ResponseWriter, r *http.Request, b *bolt.Bucket) {
	splitPath := strings.Split(r.URL.Path, "/")
	if len(splitPath) < 2 {
		log.Print("no callsign?")
		w.WriteHeader(400)
		return
	}
	callsign := strings.ToUpper(splitPath[len(splitPath)-1])
	rawRecord := b.Get([]byte(callsign))
	if rawRecord == nil {
		log.Printf("callsign does not exist: %v", callsign)
		w.WriteHeader(404)
		return
	}

	record := fcc.ReadRecord(rawRecord)
	jsonReponse, err := json.MarshalIndent(record, "", "\t")
	if err != nil {
		log.Printf("license %q parsing error: %v", callsign, err)
		w.WriteHeader(500)
		return
	}
	fmt.Fprintf(w, "%s\n", jsonReponse)
}

func main() {
	var dbPath string
	if os.Getenv("FCCDB") != "" {
		dbPath = os.Getenv("FCCDB")
	} else {
		dbPath = "artifacts/fcc.db"
	}

	options := &bolt.Options{
		ReadOnly: true,
	}

	db, err := bolt.Open(dbPath, 0444, options)
	if err != nil {
		log.Fatalf("error opening database: %s", err)
	}
	defer db.Close()

	tx, err := db.Begin(false)
	if err != nil {
		log.Fatalf("error starting transaction: %s", err)
	}
	defer tx.Rollback()

	b := tx.Bucket([]byte("licenses"))

	wrapLicenseHandler := func(w http.ResponseWriter, r *http.Request) {
		licenseHandler(w, r, b)
	}

	http.HandleFunc("/license/", wrapLicenseHandler)
	http.ListenAndServe(":8080", nil)
}
