package gonfe

import (
	"fmt"
)

const urlConsSitNFeAM = "https://nfe.sefaz.am.gov.br/services2/services/NfeConsulta4"
const urlConsSitNFeBA = "https://nfe.sefaz.ba.gov.br/webservices/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
const urlConsSitNFeCE = "https://nfe.sefaz.ce.gov.br/nfe4/services/NFeConsultaProtocolo4?wsdl"
const urlConsSitNFeGO = "https://nfe.sefaz.go.gov.br/nfe/services/NFeConsultaProtocolo4"
const urlConsSitNFeMG = "https://nfe.fazenda.mg.gov.br/nfe2/services/NFeConsultaProtocolo4"
const urlConsSitNFeMS = "https://nfe.sefaz.ms.gov.br/ws/NFeConsultaProtocolo4"
const urlConsSitNFeMT = "https://nfe.sefaz.mt.gov.br/nfews/v2/services/NfeConsulta4"
const urlConsSitNFePE = "https://nfe.sefaz.pe.gov.br/nfe-service/services/NFeConsultaProtocolo4"
const urlConsSitNFePR = "https://nfe.sefa.pr.gov.br/nfe/NFeConsultaProtocolo4"
const urlConsSitNFeRS = "https://nfe.sefazrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"
const urlConsSitNFeSP = "https://nfe.fazenda.sp.gov.br/ws/nfeconsultaprotocolo4.asmx"
const urlConsSitNFeSVAN = "https://www.sefazvirtual.fazenda.gov.br/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
const urlConsSitNFeSVRS = "https://nfe.svrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"
const urlConsSitNFeSVCAN = "https://www.svc.fazenda.gov.br/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
const urlConsSitNFeSVCRS = "https://nfe.svrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"

func getURLWS(cUF string, ws TWebService) (string, error) {
	switch ws {
	case ConsultaProtocolo:
		switch cUF {
		case "11", "12", "14", "15", "16", "17", "22", "24", "25", "27", "28", "32", "33", "42", "53":
			return urlConsSitNFeSVRS, nil
		case "13":
			return urlConsSitNFeAM, nil
		case "21":
			return urlConsSitNFeSVAN, nil
		case "23":
			return urlConsSitNFeCE, nil
		case "26":
			return urlConsSitNFePE, nil
		case "29":
			return urlConsSitNFeBA, nil
		case "31":
			return urlConsSitNFeMG, nil
		case "35":
			return urlConsSitNFeSP, nil
		case "41":
			return urlConsSitNFePR, nil
		case "43":
			return urlConsSitNFeRS, nil
		case "50":
			return urlConsSitNFeMS, nil
		case "51":
			return urlConsSitNFeMT, nil
		case "52":
			return urlConsSitNFeGO, nil
		}
	}

	return "", fmt.Errorf("WebService não encontrado: %v em %v", ws, cUF)
}
