package main

import (
	"github.com/Darklabel91/BrazilianLaws"
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"github.com/Darklabel91/Criminal_Classifier"
	"github.com/Darklabel91/Penal/csvCriminal"
)

const Path = "csvCriminal/example.csv"

var allLawsuits []csvCriminal.CriminalAnalysis

func main() {
	//read the csv
	csvFile, _ := csvCriminal.ReadCsvFile(Path, ',')

	//main
	for _, line := range csvFile {
		data, _ := classifier(line)
		allLawsuits = append(allLawsuits, data)
	}

	//export the criminal lawsuits
	csvCriminal.WriteCSV("all", "csvCriminal/Results", allLawsuits)
}

//classifier return CriminalAnalysis for one CsvStruct searched
func classifier(line csvCriminal.CsvStruct) (csvCriminal.CriminalAnalysis, error) {

	decomposedCNJ, err := CNJ.DecomposeCNJ(line.CnjNumber)

	subject := line.Subjects
	criminal, err := Criminal_Classifier.LawsuitIsCriminal(line.CnjNumber, line.Nature, line.Subjects)
	if criminal != true {
		subject = line.Subject
		criminal, err = Criminal_Classifier.LawsuitIsCriminal(line.CnjNumber, line.Nature, line.Subject)
	}

	legislation := BrazilianLaws.LawsFound(line.LawsViaCnjSubject)
	if legislation == nil {
		legislation = BrazilianLaws.SubjectLawSearch(subject)
	}
	bestLaw := BrazilianLaws.FetchBestLaw(legislation)

	return csvCriminal.CriminalAnalysis{
		CNJ:                line.CnjNumber,
		CNJYear:            decomposedCNJ.ProtocolYear,
		CNJDistrict:        decomposedCNJ.District,
		CNJUF:              decomposedCNJ.UF,
		Court:              line.Court,
		Forum:              line.Forum,
		CourtSection:       line.CourtSection,
		DocIdType:          line.DocIdType,
		DocIdNumber:        line.DocIdNumber,
		CanonicalName:      line.CanonicalName,
		CanonicalType:      line.CanonicalType,
		CoverName:          line.CoverName,
		Nature:             line.Nature,
		Subject:            line.Subject,
		Subjects:           line.Subjects,
		LawsViaCnjSubject:  line.LawsViaCnjSubject,
		Pole:               line.Pole,
		IsCriminal:         line.IsCriminal,
		IsCarta:            line.IsCarta,
		IsCriminalVerifier: criminal,
		LawArticle:         bestLaw.LawArticle,
		Law:                bestLaw.Law,
		LawNickname:        bestLaw.LawNickname,
		LawDefinition:      bestLaw.LawDefinition,
		RelatedLawsuits:    line.RelatedLawsuits,
		AllLaws:            legislation,
	}, err
}
