package main

import (
	"fmt"
	"github.com/Darklabel91/Penal/CriminalCSV"
	"github.com/Darklabel91/Penal/CriminalLaw"
	"github.com/Darklabel91/Penal/Structs"
)

const Path = "/Users/danielfillol/Downloads/CsvStruct sem título - Página1.csv"

func main() {
	//read the csv
	csvFile, err := CriminalCSV.ReadCsvFile(Path, ',')

	//main
	var criminalData []Structs.CriminalAnalysis
	var other []Structs.CriminalAnalysis
	for _, line := range csvFile {
		data := CriminalLaw.GetResult(line)
		if data.IsCriminalVerifier {
			criminalData = append(criminalData, data)
		} else {
			other = append(other, data)
		}
	}

	//export the criminal lawsuits
	err = CriminalCSV.WriteCSV("criminal", "Results", criminalData)
	if err != nil {
		fmt.Println(err)
	}

	//export every lawsuit that is not of criminal nature
	err = CriminalCSV.WriteCSV("outros", "Results", criminalData)
	if err != nil {
		fmt.Println(err)
	}
}
