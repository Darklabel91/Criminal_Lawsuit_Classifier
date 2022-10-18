package csvCriminal

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
)

//WriteCSV exports a csv to a given folder
func WriteCSV(fileName string, folderName string, Rows []CriminalAnalysis) error {
	var rows [][]string

	rows = append(rows, generateHeaders())

	for _, cnj := range Rows {
		rows = append(rows, generateRow(cnj))
	}

	cf, err := createFile(folderName + "/" + fileName + ".csv")
	if err != nil {
		return err
	}

	defer cf.Close()

	w := csv.NewWriter(cf)

	err = w.WriteAll(rows)
	if err != nil {
		return err
	}

	return nil
}

//createFile create csv file from operating system
func createFile(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

//generateHeaders generate the necessary headers for csv file
func generateHeaders() []string {
	return []string{
		"CNJ",
		"CNJYear",
		"CNJDistrict",
		"CNJUF",
		"Court",
		"Forum",
		"CourtSection",
		"DocIdType",
		"DocIdNumber",
		"CanonicalName",
		"CanonicalType",
		"CoverName",
		"NATURE",
		"SUBJECT",
		"Subjects",
		"LawsViaCnjSubject",
		"Pole",
		"IsCriminal",
		"IsCarta",
		"RelatedLawsuits",
		"JusticeSegmentVerifier",
		"NatureVerifier",
		"NatureVerifier2",
		"SubjectVerifier",
		"IsCriminalVerifier",
		"LawArticle",
		"Law",
		"LawNickname",
		"LawDefinition",
		"AllLawsFound",
	}
}

//generateRow returns a []string that compose the row in the csv file
func generateRow(row CriminalAnalysis) []string {
	var laws string
	for _, law := range row.AllLaws {
		laws += "{" + law.LawArticle + "," + law.Law + "," + law.LawNickname + "," + law.LawDefinition + "} "
	}

	return []string{
		row.CNJ,
		row.CNJYear,
		row.CNJDistrict,
		row.CNJUF,
		row.Court,
		row.Forum,
		row.CourtSection,
		row.DocIdType,
		row.DocIdNumber,
		row.CanonicalName,
		row.CanonicalType,
		row.CoverName,
		row.Nature,
		row.Subject,
		row.Subjects,
		row.LawsViaCnjSubject,
		row.Pole,
		row.IsCriminal,
		row.IsCarta,
		row.RelatedLawsuits,
		row.JusticeSegmentVerifier,
		row.NatureVerifier,
		row.NatureVerifierCriminalLaw,
		row.SubjectVerifier,
		strconv.FormatBool(row.IsCriminalVerifier),
		row.LawArticle,
		row.Law,
		row.LawNickname,
		row.LawDefinition,
		laws,
	}
}
