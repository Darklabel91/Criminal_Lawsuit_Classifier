package criminalLaw

import (
	"github.com/Darklabel91/Penal/csvCriminal"
	"strings"
)

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
			if law.LawDefinition != "?" && law.LawNickname == "Código Penal" || law.LawNickname == "Lei de Drogas" || law.LawNickname == "Lei de Armas" || law.LawNickname == "Lei Maria da Penha" || law.LawNickname == "Lei das Contravenções Penais" {
				return law
			} else if law.LawDefinition != "?" {
				return law
			}
		}
		return laws[0]
	}
	return csvCriminal.LawAnalysis{}
}
