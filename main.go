package main

import (
	"github.com/Darklabel91/Penal/criminalLaw"
	"github.com/Darklabel91/Penal/csvCriminal"
)

const Path = "example.csv"

var (
	criminalLawsuits    []csvCriminal.CriminalAnalysis
	nonCriminalLawsuits []csvCriminal.CriminalAnalysis
	allLawsuits         []csvCriminal.CriminalAnalysis
)

func main() {
	//read the csv
	csvFile, _ := csvCriminal.ReadCsvFile(Path, ',')

	//main
	for _, line := range csvFile {
		data, _ := criminalLaw.Classifier(line)
		lawsuitCSV(data)
	}

	//export the criminal lawsuits
	csvCriminal.Export(criminalLawsuits, nonCriminalLawsuits, allLawsuits)
}

//lawsuitCSV appends data in specific arrays
func lawsuitCSV(data csvCriminal.CriminalAnalysis) {
	allLawsuits = append(allLawsuits, data)
	if data.IsCriminalVerifier {
		criminalLawsuits = append(criminalLawsuits, data)
	} else {
		nonCriminalLawsuits = append(nonCriminalLawsuits, data)
	}
}
