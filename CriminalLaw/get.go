package CriminalLaw

import (
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"github.com/Darklabel91/Penal/Structs"
	"strings"
)

//GetResult return CriminalAnalysis for one CsvStruct searched
func GetResult(line Structs.CsvStruct) Structs.CriminalAnalysis {
	justiceSegmentVerifier := getCNJSegment(line.CnjNumber)
	natureVerifier := getNatureCivil(line.Nature)
	natureVerifierCriminalLaw := getCriminalLaw(line.Nature)
	subjectVerifier := getCriminalLaw(line.Subject)
	criminal := isCriminal(justiceSegmentVerifier, natureVerifier, natureVerifierCriminalLaw, subjectVerifier)

	decompose, _ := CNJ.DecomposeCNJ(line.CnjNumber)

	return Structs.CriminalAnalysis{
		CNJ:                       line.CnjNumber,
		CNJYear:                   decompose.ProtocolYear,
		Nature:                    line.Nature,
		Subject:                   line.Subject,
		IsCriminal:                line.IsCriminal,
		JusticeSegmentVerifier:    justiceSegmentVerifier,
		NatureVerifier:            natureVerifier,
		NatureVerifierCriminalLaw: natureVerifierCriminalLaw,
		SubjectVerifier:           subjectVerifier,
		IsCriminalVerifier:        criminal,
	}
}

//getCNJSegment returns "ok" for cnj segments that may contain criminal lawsuits (8,4,3,1,0)
func getCNJSegment(cnj string) string {
	decomposedCNJ, _ := CNJ.DecomposeCNJ(cnj)

	if decomposedCNJ.Segment == "8" || decomposedCNJ.Segment == "4" || decomposedCNJ.Segment == "3" || decomposedCNJ.Segment == "1" || decomposedCNJ.Segment == "0" {
		return "ok"
	} else {
		return "no"
	}
}

//getNatureCivil returns "ok" for lawsuit nature that does not contain civil
func getNatureCivil(searchWord string) string {
	civil := []string{
		"cível",
		"civel",
	}

	if searchWord == "" {
		return "vazio"
	} else {
		for _, civilWord := range civil {
			if strings.Contains(strings.ToLower(searchWord), civilWord) {
				return "no"
			}
		}
	}

	return "ok"
}

//getCriminalLaw returns "ok" if the searched word contains any match with CriminalLaw struct
func getCriminalLaw(searchWord string) string {
	if searchWord == "" {
		return "vazio"
	} else {
		if strings.Contains(searchWord, "Esbulho / Turbação / Ameaça") {
			return "no"
		}
		for _, nat := range Structs.CriminalLaw {
			if strings.Contains(strings.ToLower(searchWord), nat) {
				return "ok"
			}
		}
	}

	return "no"
}

//isCriminal returns true for the combine parameters that match with criminal law
func isCriminal(JusticeType string, CivilNature string, CriminalNature string, CriminalSubject string) bool {
	if JusticeType == "ok" && CivilNature == "ok" && CriminalNature == "ok" && CriminalSubject == "ok" {
		return true
	} else if JusticeType == "ok" && CivilNature == "ok" && CriminalNature == "ok" && CriminalSubject == "vazio" {
		return true
	} else if JusticeType == "ok" && CivilNature == "ok" && CriminalSubject == "ok" {
		return true
	} else if JusticeType == "ok" && CivilNature == "ok" && CriminalNature == "ok" && CriminalSubject != "no" {
		return true
	} else if JusticeType == "ok" && CriminalSubject == "ok" {
		return true
	} else if JusticeType == "ok" && CivilNature == "ok" && CriminalNature == "ok" {
		return true
	} else {
		return false
	}
}
