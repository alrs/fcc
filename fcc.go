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

const Sep = "\u001e"

type MinimalLicense struct {
	Callsign string
	Name     string
	Address  string
	City     string
	State    string
	ZIP      string
}

type License struct {
	LicenseID          uint32
	SourceSystem       string
	Callsign           string
	FacilityID         *uint32
	FRN                *uint32
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
	// LicAttentionLine   string
	//	ContactCompany       string
	//	ContactName          string
	//	ContactTitle         string
	//	ContactAddress1      string
	//	ContactAddress2      string
	//	ContactCity          string
	//	ContactState         string
	//	ContactZip           string
	//	ContactCountry       string
	//	ContactPhone         string
	//	ContactFax           string
	//	ContactEMail         string
	//	MarketCode           string
	//	MarketDesc           string
	//	ChannelBlock         string
	//	LocTypeCode          string
	//	LocTypeDesc          string
	//	LocCity              string
	//	LocCountyCode        string
	//	LocCountyName        string
	//	LocState             string
	//	LocRadiusOp          string
	//	LocSeqID             string
	//	LocLatDeg            string
	//	LocLatMin            string
	//	LocLatSec            string
	//	LocLatDir            string
	//	LocLongDeg           string
	//	LocLongMin           string
	//	LocLongSec           string
	//	LocLongDir           string
	//	HGTStructure         string
	//	ASRNum               string
	//	AntennaID            string
	//	AntSeqID             string
	//	AntMake              string
	//	AntModel             string
	//	AntTypeCode          string
	//	AntTypeDesc          string
	//	Azimuth              string
	//	Beamwidth            string
	//	PolarizationCode     string
	//	FrequencyID          string
	//	FreqSeqID            string
	//	FreqClassStationCode string
	//	FreqClassStationDesc string
	//	PowerERP             string
	//	PowerOutput          string
	//	FrequencyAssigned    string
	//	FrequencyUpperBand   string
	//	UnitOfMeasure        string
	//	Tolerance            string
	//	EmissionID           string
	//	EmissionSeqID        string
	//	EmissionCode         string
	//	GroundElevation      string
}
