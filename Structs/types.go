package Structs

type CriminalAnalysis struct {
	CNJ                       string `json:"CNJ,omitempty"`
	Nature                    string `json:"Nature,omitempty"`
	Subject                   string `json:"Subject,omitempty"`
	IsCriminal                string `json:"IsCriminal,omitempty"`
	JusticeSegmentVerifier    string `json:"JusticeSegmentVerifier,omitempty"`
	NatureVerifier            string `json:"NatureVerifier,omitempty"`
	NatureVerifierCriminalLaw string `json:"NatureVerifierCriminalLaw,omitempty"`
	SubjectVerifier           string `json:"SubjectVerifier,omitempty"`
	IsCriminalVerifier        bool   `json:"IsCriminalVerifier,omitempty"`
}

type CsvStruct struct {
	CnjNumber         string `json:"CnjNumber,omitempty"`
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
	Pole              string `json:"Pole,omitempty"`
	IsCriminal        string `json:"IsCriminal,omitempty"`
	IsCarta           string `json:"IsCarta,omitempty"`
}
