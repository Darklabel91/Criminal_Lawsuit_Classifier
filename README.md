# Criminal Lawsuit Classifier
Simple way of reading a .csv file to determine if a lawsuit is of criminal nature in Brazilian justice system.


## Dependecies
- [CNJ_Validate](https://github.com/Darklabel91/CNJ_Validate)
- [Criminal_Classifier](https://github.com/Darklabel91/Criminal_Classifier)

## CSV 

### Reading
.csv must have the columns as follows:
``` 
//CsvStruct csv file send to process
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
```

### Writing
.csv return as follows:
``` 
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
	AllLaws                   []LawAnalysis
}

```


## Run
Just go run main.go after chaging the constant path for the desire .csv

