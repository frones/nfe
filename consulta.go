package gonfe

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const verConsSitNFe = "4.00"
const xmlnsConsSitNFe = "http://www.portalfiscal.inf.br/nfe/wsdl/NFeConsultaProtocolo4"
const urlConsSitNFeSP = "https://nfe.fazenda.sp.gov.br/ws/nfeconsultaprotocolo4.asmx"

type ConsSitNFe struct {
	XMLName xml.Name `json:"-" xml:"http://www.portalfiscal.inf.br/nfe consSitNFe"`
	Versao  string   `json:"-" xml:"versao,attr"`
	TpAmb   TAmb     `json:"tpAmb" xml:"tpAmb"`
	XServ   string   `json:"-" xml:"xServ"`
	ChNFe   string   `json:"chNFe" xml:"chNFe"`
}

type RetConsSitNFe struct {
	XMLName       xml.Name         `json:"-" xml:"http://www.portalfiscal.inf.br/nfe retConsSitNFe"`
	Versao        string           `json:"-" xml:"versao,attr"`
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

func (c ConsSitNFe) Consulta(certFile string, certKeyFile string) (RetConsSitNFe, []byte, error) {
	if c.Versao == "" {
		c.Versao = verConsSitNFe
	}
	if c.XServ == "" {
		c.XServ = "CONSULTAR"
	}

	xmlfile, err := xml.Marshal(c)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	xmlfile, err = GetSoapEnvelope(xmlfile, xmlnsConsSitNFe)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	cert, err := tls.LoadX509KeyPair(certFile, certKeyFile)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	client := http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateFreelyAsClient,
			},
		},
	}
	req, err := http.NewRequest("POST", urlConsSitNFeSP, bytes.NewBuffer(xmlfile))
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")

	resp, err := client.Do(req)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return RetConsSitNFe{}, nil, fmt.Errorf("Falha na consulta à receita: %v", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	xmlfile, err = ReadSoapEnvelope(body)
	if err != nil {
		return RetConsSitNFe{}, nil, err
	}

	var ret RetConsSitNFe
	err = xml.Unmarshal(xmlfile, &ret)

	return ret, xmlfile, err
}
