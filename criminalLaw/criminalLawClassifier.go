package criminalLaw

import (
	"github.com/Darklabel91/Penal/csvCriminal"
)

//Classifier return CriminalAnalysis for one CsvStruct searched
func Classifier(line csvCriminal.CsvStruct) (csvCriminal.CriminalAnalysis, error) {
	cnj, err := getCNJAnalysis(line.CnjNumber)
	natureVerifier := getNatureCivil(line.Nature)
	natureVerifierCriminalLaw := getCriminalLaw(line.Nature)
	subjectVerifier := getCriminalLaw(line.Subject)
	criminal := isCriminal(cnj.Status, natureVerifier, natureVerifierCriminalLaw, subjectVerifier)
	legislation := lawsFound(line.LawsViaCnjSubject)
	bestLaw := fetchBestLaw(legislation)

	return csvCriminal.CriminalAnalysis{
		CNJ:                       line.CnjNumber,
		CNJYear:                   cnj.Year,
		CNJDistrict:               cnj.District,
		CNJUF:                     cnj.Uf,
		Court:                     line.Court,
		Forum:                     line.Forum,
		CourtSection:              line.CourtSection,
		DocIdType:                 line.DocIdType,
		DocIdNumber:               line.DocIdNumber,
		CanonicalName:             line.CanonicalName,
		CanonicalType:             line.CanonicalType,
		CoverName:                 line.CoverName,
		Nature:                    line.Nature,
		Subject:                   line.Subject,
		LawsViaCnjSubject:         line.LawsViaCnjSubject,
		Pole:                      line.Pole,
		IsCriminal:                line.IsCriminal,
		IsCarta:                   line.IsCarta,
		JusticeSegmentVerifier:    cnj.Status,
		NatureVerifier:            natureVerifier,
		NatureVerifierCriminalLaw: natureVerifierCriminalLaw,
		SubjectVerifier:           subjectVerifier,
		IsCriminalVerifier:        criminal,
		LawArticle:                bestLaw.LawArticle,
		Law:                       bestLaw.Law,
		LawNickname:               bestLaw.LawNickname,
		LawDefinition:             bestLaw.LawDefinition,
		AllLaws:                   legislation,
	}, err
}
