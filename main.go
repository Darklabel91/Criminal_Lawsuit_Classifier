package main

import (
	"github.com/Darklabel91/Penal/criminalLaw"
	"github.com/Darklabel91/Penal/csvCriminal"
)

const Path = "example.csv"

var allLawsuits []csvCriminal.CriminalAnalysis

func main() {
	//read the csv
	csvFile, _ := csvCriminal.ReadCsvFile(Path, ',')

	//main
	for _, line := range csvFile {
		data, _ := criminalLaw.Classifier(line)
		allLawsuits = append(allLawsuits, data)
	}

	//export the criminal lawsuits
	csvCriminal.WriteCSV("all", "csvCriminal/Results", allLawsuits)
}
