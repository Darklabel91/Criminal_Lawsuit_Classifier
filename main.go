package main

import (
	"fmt"
	"github.com/Darklabel91/Penal/CriminalCSV"
	"github.com/Darklabel91/Penal/CriminalLaw"
	"github.com/Darklabel91/Penal/Structs"
)

const Path = "/Users/danielfillol/Desktop/teste.csv"

func main() {
	//read the csv
	csvFile, err := CriminalCSV.ReadCsvFile(Path, ',')

	//main
	var criminalData []Structs.CriminalAnalysis
	var other []Structs.CriminalAnalysis
	var all []Structs.CriminalAnalysis
	for _, line := range csvFile {
		data := CriminalLaw.GetResult(line)

		all = append(all, data)

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
	err = CriminalCSV.WriteCSV("outros", "Results", other)
	if err != nil {
		fmt.Println(err)
	}

	//export all data
	err = CriminalCSV.WriteCSV("all", "Results", all)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Total:", len(all))
	fmt.Println("Criminal:", len(criminalData))
	fmt.Println("Outros:", len(other))
}
