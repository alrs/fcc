package fcc

func PopulateStation(line []string) *License {
	dl := License{}
	dl.LicenseID = line[0]
	dl.SourceSystem = line[1]
	dl.Callsign = line[2]
	dl.FacilityID = line[3]
	dl.FRN = line[4]
	dl.LicName = line[5]
	dl.CommonName = line[6]
	dl.RadioServiceCode = line[7]
	dl.RadioServiceDesc = line[8]
	dl.RollupCategoryCode = line[9]
	dl.RollupCategoryDesc = line[10]
	dl.GrantDate = line[11]
	dl.ExpiredDate = line[12]
	dl.CancellationDate = line[13]
	dl.LastActionDate = line[14]
	dl.LicStatusCode = line[15]
	dl.LicStatusDesc = line[16]
	dl.RollupStatusCode = line[17]
	dl.RollupStatusDesc = line[18]
	dl.EntityTypeCode = line[19]
	dl.EntityTypeDesc = line[20]
	dl.RollupEntityCode = line[21]
	dl.RollupEntityDesc = line[22]
	dl.LicAddress = line[23]
	dl.LicCity = line[24]
	dl.LicState = line[25]
	dl.LicZipCode = line[26]
	dl.LicAttentionLine = line[27]

	return &dl
}
