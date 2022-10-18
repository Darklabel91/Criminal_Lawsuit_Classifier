package csvCriminal

import (
	"encoding/csv"
	"os"
)

//ReadCsvFile reads a csv file from a given path
func ReadCsvFile(filePath string, separator rune) ([]CsvStruct, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer csvFile.Close()

	csvR := csv.NewReader(csvFile)
	csvR.Comma = separator

	csvData, err := csvR.ReadAll()
	if err != nil {
		return nil, err
	}

	var data []CsvStruct
	for i, line := range csvData {
		if i != 0 {
			var newData CsvStruct
			newData.CnjNumber = line[0]
			newData.DistributionYear = line[1]
			newData.DocIdType = line[2]
			newData.DocIdNumber = line[3]
			newData.CanonicalName = line[4]
			newData.CanonicalType = line[5]
			newData.CoverName = line[6]
			newData.Court = line[7]
			newData.Forum = line[8]
			newData.CourtSection = line[9]
			newData.Nature = line[10]
			newData.Subject = line[11]
			newData.LawsViaCnjSubject = line[12]
			newData.Subjects = line[13]
			newData.Pole = line[14]
			newData.IsCriminal = line[15]
			newData.IsCarta = line[16]
			newData.RelatedLawsuits = line[17]
			data = append(data, newData)
		}

	}

	return data, nil
}
