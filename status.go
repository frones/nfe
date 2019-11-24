package nfe

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"time"
)

const VerConsStatServ = "4.00"
const xmlnsConsStatServ = "http://www.portalfiscal.inf.br/nfe/wsdl/NFeStatusServico4"
const soapActionConsStatServ = "http://www.portalfiscal.inf.br/nfe/wsdl/NFeStatusServico4/nfeStatusServicoNF"

// ConsStatServ representa o XML de consulta do status do serviço
type ConsStatServ struct {
	XMLName xml.Name `json:"-" xml:"http://www.portalfiscal.inf.br/nfe consStatServ"`
	Versao  string   `json:"versao" xml:"versao,attr"`
	TpAmb   TAmb     `json:"tpAmb" xml:"tpAmb"`
	CUF     int      `json:"cUF" xml:"cUF"`
	XServ   string   `json:"xServ" xml:"xServ"`
}

// RetConsStatServ representa o XML de retorno da Sefaz à consulta do status do serviço
type RetConsStatServ struct {
	XMLName   xml.Name  `json:"-" xml:"http://www.portalfiscal.inf.br/nfe retConsStatServ"`
	Versao    string    `json:"versao" xml:"versao,attr"`
	TpAmb     TAmb      `json:"tpAmb" xml:"tpAmb"`
	VerAplic  string    `json:"verAplic" xml:"verAplic"`
	CStat     int       `json:"cStat" xml:"cStat"`
	XMotivo   string    `json:"xMotivo" xml:"xMotivo"`
	CUF       int       `json:"cUF" xml:"cUF"`
	DhRecbto  time.Time `json:"dhRecbto" xml:"dhRecbto"`
	TMed      int       `json:"tMed" xml:"tMed"`
	DhRetorno time.Time `json:"dhRetorno,omitempty" xml:"dhRetorno,omitempty"`
	XObs      string    `json:"xObs,omitempty" xml:"xObs,omitempty"`
}

// Realiza a consulta na Sefaz correspondente (determinada automaticamente pelo cUF), utilizando o http.Client (ver NewHTTPClient) e as funções de personalização da http.Request fornecidos.
//
// Ver ConsultaStatServ() para uma maneira mais simples de consultar o status do serviço
func (cons ConsStatServ) Consulta(client *http.Client, optReq ...func(req *http.Request)) (RetConsStatServ, []byte, error) {
	url, err := getURLWS(cons.CUF, cons.TpAmb, ConsultaStatus)
	if err != nil {
		return RetConsStatServ{}, nil, err
	}

	xmlfile, err := sendRequest(cons, url, xmlnsConsStatServ, soapActionConsStatServ, client, optReq...)
	if err != nil {
		return RetConsStatServ{}, nil, fmt.Errorf("Erro na comunicação com a Sefaz. Detalhes: %w", err)
	}

	var ret RetConsStatServ
	err = xml.Unmarshal(xmlfile, &ret)
	if err != nil {
		return RetConsStatServ{}, xmlfile, fmt.Errorf("Erro na desserialização do arquivo XML: %w. Arquivo: %s", err, xmlfile)
	}

	return ret, xmlfile, nil
}

// Função auxiliar para executar a ConsStatServ.Consulta()
func ConsultaStatServ(cUF int, tpAmb TAmb, client *http.Client, optReq ...func(req *http.Request)) (RetConsStatServ, []byte, error) {
	cons := ConsStatServ{
		Versao: VerConsStatServ,
		TpAmb:  tpAmb,
		XServ:  "STATUS",
		CUF:    cUF,
	}

	return cons.Consulta(client, optReq...)
}
