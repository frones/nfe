package gonfe

import (
	"fmt"
)

const (
	urlConsSitNFeAM    = "https://nfe.sefaz.am.gov.br/services2/services/NfeConsulta4"
	urlConsSitNFeBA    = "https://nfe.sefaz.ba.gov.br/webservices/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
	urlConsSitNFeCE    = "https://nfe.sefaz.ce.gov.br/nfe4/services/NFeConsultaProtocolo4"
	urlConsSitNFeGO    = "https://nfe.sefaz.go.gov.br/nfe/services/NFeConsultaProtocolo4"
	urlConsSitNFeMG    = "https://nfe.fazenda.mg.gov.br/nfe2/services/NFeConsultaProtocolo4"
	urlConsSitNFeMS    = "https://nfe.sefaz.ms.gov.br/ws/NFeConsultaProtocolo4"
	urlConsSitNFeMT    = "https://nfe.sefaz.mt.gov.br/nfews/v2/services/NfeConsulta4"
	urlConsSitNFePE    = "https://nfe.sefaz.pe.gov.br/nfe-service/services/NFeConsultaProtocolo4"
	urlConsSitNFePR    = "https://nfe.sefa.pr.gov.br/nfe/NFeConsultaProtocolo4"
	urlConsSitNFeRS    = "https://nfe.sefazrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"
	urlConsSitNFeSP    = "https://nfe.fazenda.sp.gov.br/ws/nfeconsultaprotocolo4.asmx"
	urlConsSitNFeSVAN  = "https://www.sefazvirtual.fazenda.gov.br/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
	urlConsSitNFeSVRS  = "https://nfe.svrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"
	urlConsSitNFeSVCAN = "https://www.svc.fazenda.gov.br/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
	urlConsSitNFeSVCRS = "https://nfe.svrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"

	urlHomConsSitNFeAM    = "https://homnfe.sefaz.am.gov.br/services2/services/NfeConsulta4"
	urlHomConsSitNFeBA    = "https://hnfe.sefaz.ba.gov.br/webservices/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
	urlHomConsSitNFeCE    = "https://nfeh.sefaz.ce.gov.br/nfe4/services/NFeConsultaProtocolo4"
	urlHomConsSitNFeGO    = "https://homolog.sefaz.go.gov.br/nfe/services/NFeConsultaProtocolo4"
	urlHomConsSitNFeMG    = "https://hnfe.fazenda.mg.gov.br/nfe2/services/NFeConsultaProtocolo4"
	urlHomConsSitNFeMS    = "https://hom.nfe.sefaz.ms.gov.br/ws/NFeConsultaProtocolo4"
	urlHomConsSitNFeMT    = "https://homologacao.sefaz.mt.gov.br/nfews/v2/services/NfeConsulta4"
	urlHomConsSitNFePE    = "https://nfehomolog.sefaz.pe.gov.br/nfe-service/services/NFeConsultaProtocolo4"
	urlHomConsSitNFePR    = "https://homologacao.nfe.sefa.pr.gov.br/nfe/NFeConsultaProtocolo4"
	urlHomConsSitNFeRS    = "https://nfe-homologacao.sefazrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"
	urlHomConsSitNFeSP    = "https://homologacao.nfe.fazenda.sp.gov.br/ws/nfeconsultaprotocolo4.asmx"
	urlHomConsSitNFeSVAN  = "https://hom.sefazvirtual.fazenda.gov.br/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
	urlHomConsSitNFeSVRS  = "https://nfe-homologacao.svrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"
	urlHomConsSitNFeSVCAN = "https://hom.svc.fazenda.gov.br/NFeConsultaProtocolo4/NFeConsultaProtocolo4.asmx"
	urlHomConsSitNFeSVCRS = "https://nfe-homologacao.svrs.rs.gov.br/ws/NfeConsulta/NfeConsulta4.asmx"
)

const (
	urlConsStatServAM    = "https://nfe.sefaz.am.gov.br/services2/services/NfeStatusServico4"
	urlConsStatServBA    = "https://nfe.sefaz.ba.gov.br/webservices/NFeStatusServico4/NFeStatusServico4.asmx"
	urlConsStatServCE    = "https://nfe.sefaz.ce.gov.br/nfe4/services/NFeStatusServico4"
	urlConsStatServGO    = "https://nfe.sefaz.go.gov.br/nfe/services/NFeStatusServico4"
	urlConsStatServMG    = "https://nfe.fazenda.mg.gov.br/nfe2/services/NFeStatusServico4"
	urlConsStatServMS    = "https://nfe.sefaz.ms.gov.br/ws/NFeStatusServico4"
	urlConsStatServMT    = "https://nfe.sefaz.mt.gov.br/nfews/v2/services/NfeStatusServico4"
	urlConsStatServPE    = "https://nfe.sefaz.pe.gov.br/nfe-service/services/NFeStatusServico4"
	urlConsStatServPR    = "https://nfe.sefa.pr.gov.br/nfe/NFeStatusServico4"
	urlConsStatServRS    = "https://nfe.sefazrs.rs.gov.br/ws/NfeStatusServico/NfeStatusServico4.asmx"
	urlConsStatServSP    = "https://nfe.fazenda.sp.gov.br/ws/nfestatusservico4.asmx"
	urlConsStatServSVAN  = "https://www.sefazvirtual.fazenda.gov.br/NFeStatusServico4/NFeStatusServico4.asmx"
	urlConsStatServSVRS  = "https://nfe.svrs.rs.gov.br/ws/NfeStatusServico/NfeStatusServico4.asmx"
	urlConsStatServSVCAN = "https://www.svc.fazenda.gov.br/NFeStatusServico4/NFeStatusServico4.asmx"
	urlConsStatServSVCRS = "https://nfe.svrs.rs.gov.br/ws/NfeStatusServico/NfeStatusServico4.asmx"

	urlHomConsStatServAM    = "https://homnfe.sefaz.am.gov.br/services2/services/NfeStatusServico4"
	urlHomConsStatServBA    = "https://hnfe.sefaz.ba.gov.br/webservices/NFeStatusServico4/NFeStatusServico4.asmx"
	urlHomConsStatServCE    = "https://nfeh.sefaz.ce.gov.br/nfe4/services/NFeStatusServico4"
	urlHomConsStatServGO    = "https://homolog.sefaz.go.gov.br/nfe/services/NFeStatusServico4"
	urlHomConsStatServMG    = "https://hnfe.fazenda.mg.gov.br/nfe2/services/NFeStatusServico4"
	urlHomConsStatServMS    = "https://hom.nfe.sefaz.ms.gov.br/ws/NFeStatusServico4"
	urlHomConsStatServMT    = "https://homologacao.sefaz.mt.gov.br/nfews/v2/services/NfeStatusServico4"
	urlHomConsStatServPE    = "https://nfehomolog.sefaz.pe.gov.br/nfe-service/services/NFeStatusServico4"
	urlHomConsStatServPR    = "https://homologacao.nfe.sefa.pr.gov.br/nfe/NFeStatusServico4"
	urlHomConsStatServRS    = "https://nfe-homologacao.sefazrs.rs.gov.br/ws/NfeStatusServico/NfeStatusServico4.asmx"
	urlHomConsStatServSP    = "https://homologacao.nfe.fazenda.sp.gov.br/ws/nfestatusservico4.asmx"
	urlHomConsStatServSVAN  = "https://hom.sefazvirtual.fazenda.gov.br/NFeStatusServico4/NFeStatusServico4.asmx"
	urlHomConsStatServSVRS  = "https://nfe-homologacao.svrs.rs.gov.br/ws/NfeStatusServico/NfeStatusServico4.asmx"
	urlHomConsStatServSVCAN = "https://hom.svc.fazenda.gov.br/NFeStatusServico4/NFeStatusServico4.asmx"
	urlHomConsStatServSVCRS = "https://nfe-homologacao.svrs.rs.gov.br/ws/NfeStatusServico/NfeStatusServico4.asmx"
)

// getURLWS obtem a URL para o serviço e a UF informados.
func getURLWS(cUF int, tpAmb TAmb, ws TWebService) (string, error) {
	switch tpAmb {
	case Producao:
		switch ws {
		case ConsultaProtocolo:
			switch cUF {
			case 11, 12, 14, 15, 16, 17, 22, 24, 25, 27, 28, 32, 33, 42, 53:
				return urlConsSitNFeSVRS, nil
			case 13:
				return urlConsSitNFeAM, nil
			case 21:
				return urlConsSitNFeSVAN, nil
			case 23:
				return urlConsSitNFeCE, nil
			case 26:
				return urlConsSitNFePE, nil
			case 29:
				return urlConsSitNFeBA, nil
			case 31:
				return urlConsSitNFeMG, nil
			case 35:
				return urlConsSitNFeSP, nil
			case 41:
				return urlConsSitNFePR, nil
			case 43:
				return urlConsSitNFeRS, nil
			case 50:
				return urlConsSitNFeMS, nil
			case 51:
				return urlConsSitNFeMT, nil
			case 52:
				return urlConsSitNFeGO, nil
			}
		case ConsultaStatus:
			switch cUF {
			case 11, 12, 14, 15, 16, 17, 22, 24, 25, 27, 28, 32, 33, 42, 53:
				return urlConsStatServSVRS, nil
			case 13:
				return urlConsStatServAM, nil
			case 21:
				return urlConsStatServSVAN, nil
			case 23:
				return urlConsStatServCE, nil
			case 26:
				return urlConsStatServPE, nil
			case 29:
				return urlConsStatServBA, nil
			case 31:
				return urlConsStatServMG, nil
			case 35:
				return urlConsStatServSP, nil
			case 41:
				return urlConsStatServPR, nil
			case 43:
				return urlConsStatServRS, nil
			case 50:
				return urlConsStatServMS, nil
			case 51:
				return urlConsStatServMT, nil
			case 52:
				return urlConsStatServGO, nil
			}
		}
	case Homologacao:
		switch ws {
		case ConsultaProtocolo:
			switch cUF {
			case 11, 12, 14, 15, 16, 17, 22, 24, 25, 27, 28, 32, 33, 42, 53:
				return urlHomConsSitNFeSVRS, nil
			case 13:
				return urlHomConsSitNFeAM, nil
			case 21:
				return urlHomConsSitNFeSVAN, nil
			case 23:
				return urlHomConsSitNFeCE, nil
			case 26:
				return urlHomConsSitNFePE, nil
			case 29:
				return urlHomConsSitNFeBA, nil
			case 31:
				return urlHomConsSitNFeMG, nil
			case 35:
				return urlHomConsSitNFeSP, nil
			case 41:
				return urlHomConsSitNFePR, nil
			case 43:
				return urlHomConsSitNFeRS, nil
			case 50:
				return urlHomConsSitNFeMS, nil
			case 51:
				return urlHomConsSitNFeMT, nil
			case 52:
				return urlHomConsSitNFeGO, nil
			}
		case ConsultaStatus:
			switch cUF {
			case 11, 12, 14, 15, 16, 17, 22, 24, 25, 27, 28, 32, 33, 42, 53:
				return urlHomConsStatServSVRS, nil
			case 13:
				return urlHomConsStatServAM, nil
			case 21:
				return urlHomConsStatServSVAN, nil
			case 23:
				return urlHomConsStatServCE, nil
			case 26:
				return urlHomConsStatServPE, nil
			case 29:
				return urlHomConsStatServBA, nil
			case 31:
				return urlHomConsStatServMG, nil
			case 35:
				return urlHomConsStatServSP, nil
			case 41:
				return urlHomConsStatServPR, nil
			case 43:
				return urlHomConsStatServRS, nil
			case 50:
				return urlHomConsStatServMS, nil
			case 51:
				return urlHomConsStatServMT, nil
			case 52:
				return urlHomConsStatServGO, nil
			}
		}
	}

	return "", fmt.Errorf("WebService não encontrado: %v em %v", ws, cUF)
}
