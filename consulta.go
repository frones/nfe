package gonfe

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const verConsSitNFe = "4.00"
const xmlnsConsSitNFe = "http://www.portalfiscal.inf.br/nfe/wsdl/NFeConsultaProtocolo4"

// ConsSitNFe representa o XML de consulta de uma NFe
type ConsSitNFe struct {
	XMLName xml.Name `json:"-" xml:"http://www.portalfiscal.inf.br/nfe consSitNFe"`
	Versao  string   `json:"versao" xml:"versao,attr"`
	TpAmb   TAmb     `json:"tpAmb" xml:"tpAmb"`
	XServ   string   `json:"xServ" xml:"xServ"`
	ChNFe   string   `json:"chNFe" xml:"chNFe"`
}

// RetConsSitNFe representa o XML de retorno da Sefaz à consulta da NFe
type RetConsSitNFe struct {
	XMLName       xml.Name         `json:"-" xml:"http://www.portalfiscal.inf.br/nfe retConsSitNFe"`
	Versao        string           `json:"versao" xml:"versao,attr"`
	TpAmb         TAmb             `json:"tpAmb" xml:"tpAmb"`
	VerAplic      string           `json:"verAplic" xml:"verAplic"`
	CStat         int              `json:"cStat" xml:"cStat"`
	XMotivo       string           `json:"xMotivo" xml:"xMotivo"`
	CUF           int              `json:"cUF" xml:"cUF"`
	DhRecbto      time.Time        `json:"dhRecbto" xml:"dhRecbto"`
	ChNFe         string           `json:"chNFe" xml:"chNFe"`
	ProtNFe       *ProtNFe         `json:"protNFe" xml:"protNFe"`
	RetCancNFe    *RetCancNFe      `json:"retCancNFe,omitempty" xml:"retCancNFe,omitempty"`
	ProcEventoNFe *[]ProcEventoNFe `json:"procEventoNFe,omitempty" xml:"procEventoNFe,omitempty"`
}

// Realiza a consulta na Sefaz correspondente (determinada automaticamente pelo cUF presente na chave), utilizando o http.Client (ver NewHTTPClient) e as funções de personalização da http.Request fornecidos.
//
// Ver ConsultaNFe() para uma maneira mais simples de consultar a NFe
func (cons ConsSitNFe) Consulta(client *http.Client, optReq ...func(req *http.Request)) (RetConsSitNFe, []byte, error) {
	cUF, _, _, _, _, _, _, _, _, err := GetChaveInfo(cons.ChNFe)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}
	url, err := getURLWS(cUF, ConsultaProtocolo)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	xmlfile, err := xml.Marshal(cons)
	if err != nil {
		return RetConsSitNFe{}, nil, fmt.Errorf("Erro na geração do XML de consulta. Detalhes: %v", err)
	}

	xmlfile, err = getSoapEnvelope(xmlfile, xmlnsConsSitNFe)
	if err != nil {
		return RetConsSitNFe{}, nil, fmt.Errorf("Erro na geração do envelope SOAP. Detalhes: %v", err)
	}
	xmlfile = []byte(append([]byte(xml.Header), xmlfile...))

	req, err := newRequest(url, xmlfile)
	if err != nil {
		return RetConsSitNFe{}, nil, fmt.Errorf("Erro na criação da requisição (http.Request) para a URL %s. Detalhes: %v", url, err)
	}
	for _, opt := range optReq {
		opt(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return RetConsSitNFe{}, nil, fmt.Errorf("Erro na requisição ao WebService %s. Detalhes: %v", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return RetConsSitNFe{}, nil, fmt.Errorf("Falha na consulta à receita (%s): %v", url, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	xmlfile, err = readSoapEnvelope(body)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	var ret RetConsSitNFe
	err = xml.Unmarshal(xmlfile, &ret)

	return ret, xmlfile, err
}

// Função auxiliar para executar a ConsSitNFe.Consulta()
func ConsultaNFe(dfechave string, tpAmb TAmb, client *http.Client, optReq ...func(req *http.Request)) (RetConsSitNFe, []byte, error) {
	cons := ConsSitNFe{
		Versao: verConsSitNFe,
		TpAmb:  tpAmb,
		XServ:  "CONSULTAR",
		ChNFe:  dfechave,
	}

	return cons.Consulta(client, optReq...)
}
