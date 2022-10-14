package criminalLaw

import (
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"github.com/Darklabel91/Penal/csvCriminal"
	"strings"
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

//lawsFound search for every legislation on mapLaws
func lawsFound(law string) []csvCriminal.LawAnalysis {
	var localAllLaws []csvCriminal.LawAnalysis
	if law != "" {
		splitLaw := strings.Split(law, "-")
		if len(splitLaw) > 0 {
			for i := 0; i < len(splitLaw); i++ {
				newLaw := strings.Split(splitLaw[i], "Artigo ")
				for _, law2 := range newLaw {
					clearLawDa := strings.Replace(law2, " da ", "_", -1)
					clearLaw := strings.Replace(clearLawDa, " do ", "_", -1)
					legislation, err := fetchLegislation(strings.TrimSpace(clearLaw))
					if err == nil {
						localAllLaws = append(localAllLaws, legislation)
					}
				}
			}
		}
	}

	return localAllLaws
}

//fetchBestLaw fetch the best possible law
func fetchBestLaw(laws []csvCriminal.LawAnalysis) csvCriminal.LawAnalysis {
	if len(laws) > 0 {
		for _, law := range laws {
			if law.LawDefinition != "?" && law.LawNickname == "Código Penal" || law.LawNickname == "Lei de Drogas" || law.LawNickname == "Lwi de Armas" || law.LawNickname == "Lei Maria da Penha" || law.LawNickname == "Lei das Contravenções Penais" {
				return law
			} else if law.LawDefinition != "?" {
				return law
			}
		}
		return laws[0]
	}
	return csvCriminal.LawAnalysis{}
}

//getCNJAnalysis returns cnj district and uf and the status: "ok" for cnj segments that may contain criminal lawsuits (8,4,3,1,0)
func getCNJAnalysis(cnj string) (csvCriminal.CnjAnalysis, error) {
	decomposedCNJ, err := CNJ.DecomposeCNJ(cnj)

	var status string
	if decomposedCNJ.Segment == "8" || decomposedCNJ.Segment == "4" || decomposedCNJ.Segment == "3" || decomposedCNJ.Segment == "1" || decomposedCNJ.Segment == "0" {
		status = "ok"
	} else {
		status = "no"
	}

	if err != nil {
		return csvCriminal.CnjAnalysis{
			Status:   status,
			Year:     decomposedCNJ.ProtocolYear,
			District: "",
			Uf:       "",
		}, err
	}

	return csvCriminal.CnjAnalysis{
		Status:   status,
		Year:     decomposedCNJ.ProtocolYear,
		District: decomposedCNJ.District,
		Uf:       decomposedCNJ.UF,
	}, nil
}

//getNatureCivil returns "ok" for lawsuit nature that does not contain civil
func getNatureCivil(searchWord string) string {
	civil := []string{
		"cível",
		"civel",
		"civ",
		"procedimento comum",
		"execução fiscal",
		"execução de título extrajudicial",
		"busca e apreensão",
		"monitória",
		"usucapião",
		"trabalhista",
		"reclamação pré-processual",
		"processo de conhecimento",
		"cumprimento de sentença",
	}

	if searchWord == "" {
		return "vazio"
	} else {
		for _, penalWord := range criminalLaw {
			if strings.Contains(strings.ToLower(searchWord), penalWord) {
				return "ok"
			}
		}
		for _, civilWord := range civil {
			if strings.Contains(strings.ToLower(searchWord), civilWord) {
				return "no"
			}
		}

	}

	return "no"
}

//criminalLaw struct that contains possible matches with criminal law
var criminalLaw = []string{
	"cpp",
	"ceman",
	"execução penal",
	"prestação pecuniária",
	"prestação de serviços à comunidade",
	"descaminho",
	"interrogatório",
	"atos executórios",
	"inquirição",
	"oitiva",
	"lesão corporal",
	"estelionato",
	"injúria",
	"difamação",
	"ameaça",
	"homicídio",
	"preso",
	"penal",
	"criminal",
	"inquérito",
	"inquerito",
	"pena",
	"tribunal do juri",
	"corpus",
	"liberdade",
	"reclusão",
	"prisão",
	"prisao",
	"antitóxicos",
	"crime",
	"restituição de coisas",
	"termo circunstanciado",
	"maria da penh",
	"núm.caract(g153,)<>núm.caract(substituir(g153",
	"maria da penha",
	"crim",
	"proced.investigatório do mp",
	"detenção",
	"execução de medidas alternativas",
	"toxico",
	"indulto",
	"acusado",
	"pedido de quebra de sigilo de dados e/ou telefônic",
	"semi-aberto",
	"fechado",
	"aberto",
	"condicional",
	"furto",
	"roubo",
	"receptação",
	"latrocínio",
	"drogas",
	"apropriação indébita",
	"dano qualificad",
	"falso",
	"policial",
	"violência",
}

//getCriminalLaw returns "ok" if the searched word contains any match with criminalLaw struct
func getCriminalLaw(searchWord string) string {
	if searchWord == "" {
		return "vazio"
	} else {
		if strings.Contains(searchWord, "Esbulho / Turbação / Ameaça") {
			return "no"
		}
		for _, nat := range criminalLaw {
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
