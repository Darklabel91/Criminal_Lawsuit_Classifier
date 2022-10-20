package csvCriminal

import "github.com/Darklabel91/BrazilianLaws/laws"

//CsvStruct csv file send to process
type CsvStruct struct {
	CnjNumber         string `json:"CnjNumber,omitempty"`
	DistributionYear  string `json:"DistributionYear,omitempty"`
	DocIdType         string `json:"DocIdType,omitempty"`
	DocIdNumber       string `json:"DocIdNumber,omitempty"`
	CanonicalName     string `json:"CanonicalName,omitempty"`
	CanonicalType     string `json:"CanonicalType,omitempty"`
	CoverName         string `json:"CoverName,omitempty"`
	Court             string `json:"Court,omitempty"`
	Forum             string `json:"Forum,omitempty"`
	CourtSection      string `json:"CourtSection,omitempty"`
	Nature            string `json:"Nature,omitempty"`
	Subject           string `json:"Subject,omitempty"`
	LawsViaCnjSubject string `json:"LawsViaCnjSubject,omitempty"`
	Subjects          string `json:"Subjects,omitempty"`
	Pole              string `json:"Pole,omitempty"`
	IsCriminal        string `json:"IsCriminal,omitempty"`
	IsCarta           string `json:"IsCarta,omitempty"`
	RelatedLawsuits   string `json:"RelatedLawsuits,omitempty"`
}

//CriminalAnalysis output
type CriminalAnalysis struct {
	CNJ                       string `json:"CNJ,omitempty"`
	CNJYear                   string `json:"CNJYear,omitempty"`
	CNJDistrict               string `json:"CNJDistrict,omitempty"`
	CNJUF                     string `json:"CNJUF,omitempty"`
	Court                     string `json:"Court,omitempty"`
	Forum                     string `json:"Forum,omitempty"`
	CourtSection              string `json:"CourtSection,omitempty"`
	DocIdType                 string `json:"DocIdType,omitempty"`
	DocIdNumber               string `json:"DocIdNumber,omitempty"`
	CanonicalName             string `json:"CanonicalName,omitempty"`
	CanonicalType             string `json:"CanonicalType,omitempty"`
	CoverName                 string `json:"CoverName,omitempty"`
	Nature                    string `json:"Nature,omitempty"`
	Subject                   string `json:"Subject,omitempty"`
	Subjects                  string `json:"Subjects,omitempty"`
	LawsViaCnjSubject         string `json:"LawsViaCnjSubject,omitempty"`
	Pole                      string `json:"Pole,omitempty"`
	IsCriminal                string `json:"IsCriminal,omitempty"`
	IsCarta                   string `json:"IsCarta,omitempty"`
	JusticeSegmentVerifier    string `json:"JusticeSegmentVerifier,omitempty"`
	NatureVerifier            string `json:"NatureVerifier,omitempty"`
	NatureVerifierCriminalLaw string `json:"NatureVerifierCriminalLaw,omitempty"`
	SubjectVerifier           string `json:"SubjectVerifier,omitempty"`
	IsCriminalVerifier        bool   `json:"IsCriminalVerifier,omitempty"`
	LawArticle                string `json:"LawArticle,omitempty"`
	Law                       string `json:"Law,omitempty"`
	LawNickname               string `json:"LawNickname,omitempty"`
	LawDefinition             string `json:"LawDefinition,omitempty"`
	RelatedLawsuits           string `json:"RelatedLawsuits,omitempty"`
	AllLaws                   []laws.Analysis
}
