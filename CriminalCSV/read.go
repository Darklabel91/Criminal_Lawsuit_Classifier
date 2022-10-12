package CriminalCSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Penal/Structs"
	"os"
)

//ReadCsvFile reads a csv file from a given path
func ReadCsvFile(filePath string, separator rune) ([]Structs.CsvStruct, error) {
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

	var data []Structs.CsvStruct
	for i, line := range csvData {
		if i != 0 {
			var newData Structs.CsvStruct
			newData.CnjNumber = line[0]
			newData.DocIdType = line[1]
			newData.DocIdNumber = line[2]
			newData.CanonicalName = line[3]
			newData.CanonicalType = line[4]
			newData.CoverName = line[5]
			newData.Court = line[6]
			newData.Forum = line[7]
			newData.CourtSection = line[8]
			newData.Nature = line[9]
			newData.Subject = line[10]
			newData.LawsViaCnjSubject = line[11]
			newData.Pole = line[12]
			newData.IsCriminal = line[13]
			newData.IsCarta = line[14]
			data = append(data, newData)
		}

	}

	return data, nil
}
