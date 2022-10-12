package CriminalCSV

import (
	"encoding/csv"
	"github.com/Darklabel91/Penal/Structs"
	"os"
	"path/filepath"
	"strconv"
)

//WriteCSV exports a csv to a given folder
func WriteCSV(fileName string, folderName string, Rows []Structs.CriminalAnalysis) error {
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
		"NATURE",
		"SUBJECT",
		"IsCriminal",
		"JusticeSegmentVerifier",
		"NatureVerifier",
		"NatureVerifier2",
		"SubjectVerifier",
		"IsCriminalVerifier",
	}
}

//generateRow returns a []string that compose the row in the csv file
func generateRow(row Structs.CriminalAnalysis) []string {
	return []string{
		row.CNJ,
		row.CNJYear,
		row.Nature,
		row.Subject,
		row.IsCriminal,
		row.JusticeSegmentVerifier,
		row.NatureVerifier,
		row.NatureVerifierCriminalLaw,
		row.SubjectVerifier,
		strconv.FormatBool(row.IsCriminalVerifier),
	}
}
