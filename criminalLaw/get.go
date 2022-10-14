package criminalLaw

import (
	"github.com/Darklabel91/CNJ_Validate/CNJ"
	"github.com/Darklabel91/Penal/csvCriminal"
	"strings"
)

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
