// package fcc
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

package fcc

import (
	"time"
)

// MinimalLicense is a struct that represents the subset of license
// data stored to database for fast lookups.
type MinimalLicense struct {
	Name    string
	Address string
	City    string
	State   string
	ZIP     string
}

// License is a struct representing fields of the FCC license database
// dump of interest to amateur radio operators.
type License struct {
	LicenseID          uint32
	SourceSystem       string
	Callsign           string
	FacilityID         *uint32
	FRN                *uint64
	LicName            string
	CommonName         string
	RadioServiceCode   string
	RadioServiceDesc   string
	RollupCategoryCode string
	RollupCategoryDesc string
	GrantDate          *time.Time
	ExpiredDate        *time.Time
	CancellationDate   *time.Time
	LastActionDate     *time.Time
	LicStatusCode      string
	LicStatusDesc      string
	RollupStatusCode   string
	RollupStatusDesc   string
	EntityTypeCode     string
	EntityTypeDesc     string
	RollupEntityCode   string
	RollupEntityDesc   string
	LicAddress         string
	LicCity            string
	LicState           string
	LicZipCode         string
}
