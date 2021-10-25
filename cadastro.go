package nfe

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"regexp"
	"time"

	"cloud.google.com/go/civil"
	"github.com/frones/brdocs"
)

const VerConsCad = "2.00"
const xmlnsConsCad = "http://www.portalfiscal.inf.br/nfe/wsdl/CadConsultaCadastro4"
const soapActionConsCad = "http://www.portalfiscal.inf.br/nfe/wsdl/CadConsultaCadastro4/consultaCadastro"

type InfCad struct {
	IE         string     `json:"IE" xml:"IE"`
	CNPJ       string     `json:"CNPJ,omitempty" xml:"CNPJ,omitempty"`
	CPF        string     `json:"CPF,omitempty" xml:"CPF,omitempty"`
	UF         string     `json:"UF" xml:"UF"`
	CSit       int        `json:"cSit" xml:"cSit"`
	IndCredNFe int        `json:"indCredNFe" xml:"indCredNFe"`
	IndCredCTe int        `json:"indCredCTe" xml:"indCredCTe"`
	XNome      string     `json:"xNome" xml:"xNome"`
	XFant      string     `json:"xFant,omitempty" xml:"xFant,omitempty"`
	XRegApur   string     `json:"xRegApur,omitempty" xml:"xRegApur,omitempty"`
	CNAE       string     `json:"CNAE,omitempty" xml:"CNAE,omitempty"`
	DIniAtiv   civil.Date `json:"dIniAtiv,omitempty" xml:"dIniAtiv,omitempty"`
	DUltSit    civil.Date `json:"dUltSit,omitempty" xml:"dUltSit,omitempty"`
	DBaixa     civil.Date `json:"dBaixa,omitempty" xml:"dBaixa,omitempty"`
	IEUnica    string     `json:"IEUnica,omitempty" xml:"IEUnica,omitempty"`
	IEAtual    string     `json:"IEAtual,omitempty" xml:"IEAtual,omitempty"`
	Ender      *struct {
		XLgr    string `json:"xLgr,omitempty" xml:"xLgr,omitempty"`
		Nro     string `json:"nro,omitempty" xml:"nro,omitempty"`
		XCpl    string `json:"xCpl,omitempty" xml:"xCpl,omitempty"`
		XBairro string `json:"xBairro,omitempty" xml:"xBairro,omitempty"`
		CMun    string `json:"cMun,omitempty" xml:"cMun,omitempty"`
		XMun    string `json:"xMun,omitempty" xml:"xMun,omitempty"`
		CEP     string `json:"CEP,omitempty" xml:"CEP,omitempty"`
	} `json:"ender,omitempty" xml:"ender,omitempty"`
}

// ConsCad representa o XML de consulta do cadastro do contribuinte
type ConsCad struct {
	XMLName xml.Name `json:"-" xml:"http://www.portalfiscal.inf.br/nfe ConsCad"`
	Versao  string   `json:"versao" xml:"versao,attr"`
	InfCons struct {
		XServ string `json:"xServ" xml:"xServ"`
		UF    string `json:"UF" xml:"UF"`
		IE    string `json:"IE,omitempty" xml:"IE,omitempty"`
		CNPJ  string `json:"CNPJ,omitempty" xml:"CNPJ,omitempty"`
		CPF   string `json:"CPF,omitempty" xml:"CPF,omitempty"`
	} `json:"infCons" xml:"infCons"`
}

// RetConsCad representa o XML de retorno da Sefaz à consulta do cadastro do contribuinte
type RetConsCad struct {
	XMLName xml.Name `json:"-" xml:"http://www.portalfiscal.inf.br/nfe retConsCad"`
	Versao  string   `json:"versao" xml:"versao,attr"`
	InfCons struct {
		VerAplic string    `json:"verAplic" xml:"verAplic"`
		CStat    int       `json:"cStat" xml:"cStat"`
		XMotivo  string    `json:"xMotivo" xml:"xMotivo"`
		UF       string    `json:"UF" xml:"UF"`
		IE       string    `json:"IE,omitempty" xml:"IE,omitempty"`
		CNPJ     string    `json:"CNPJ,omitempty" xml:"CNPJ,omitempty"`
		CPF      string    `json:"CPF,omitempty" xml:"CPF,omitempty"`
		DhCons   time.Time `json:"dhCons" xml:"dhCons"`
		CUF      int       `json:"cUF" xml:"cUF"`
		InfCad   *[]InfCad `json:"infCad,omitempty" xml:"infCad,omitempty"`
	} `json:"infCons" xml:"infCons"`
}

// Realiza a consulta na Sefaz correspondente (determinada automaticamente pelo UF), utilizando o http.Client (ver NewHTTPClient) e as funções de personalização da http.Request fornecidos.
//
// Ver ConsultaCad() para uma maneira mais simples de consultar o status do serviço
func (cons ConsCad) Consulta(tpAmb TAmb, client *http.Client, optReq ...func(req *http.Request)) (RetConsCad, []byte, error) {
	if cons.InfCons.IE != "" {
		if (cons.InfCons.CNPJ != "") || (cons.InfCons.CPF != "") {
			return RetConsCad{}, nil, fmt.Errorf("Erro na consulta de cadastro: apenas um documento deve ser informado (IE, CNPJ ou CPF)")
		}
		if !brdocs.ValidaIE(cons.InfCons.IE, cons.InfCons.UF) {
			return RetConsCad{}, nil, fmt.Errorf("Erro na consulta de cadastro: IE inválida: %s", cons.InfCons.IE)
		}
	}
	if cons.InfCons.CNPJ != "" {
		if (cons.InfCons.IE != "") || (cons.InfCons.CPF != "") {
			return RetConsCad{}, nil, fmt.Errorf("Erro na consulta de cadastro: apenas um documento deve ser informado (IE, CNPJ ou CPF)")
		}
		if !brdocs.ValidaCNPJ(cons.InfCons.CNPJ) {
			return RetConsCad{}, nil, fmt.Errorf("Erro na consulta de cadastro: CNPJ inválido: %s", cons.InfCons.CNPJ)
		}
	}
	if cons.InfCons.CPF != "" {
		if (cons.InfCons.CNPJ != "") || (cons.InfCons.IE != "") {
			return RetConsCad{}, nil, fmt.Errorf("Erro na consulta de cadastro: apenas um documento deve ser informado (IE, CNPJ ou CPF)")
		}
		if !brdocs.ValidaCPF(cons.InfCons.CPF) {
			return RetConsCad{}, nil, fmt.Errorf("Erro na consulta de cadastro: CPF inválido: %s", cons.InfCons.CPF)
		}
	}

	url, err := getURLWS(GetcUF(cons.InfCons.UF), tpAmb, ConsultaCadastro)
	if err != nil {
		return RetConsCad{}, nil, err
	}

	xmlfile, err := sendRequest(cons, url, xmlnsConsCad, soapActionConsCad, client, optReq...)
	if err != nil {
		return RetConsCad{}, nil, fmt.Errorf("Erro na comunicação com a Sefaz. Detalhes: %w", err)
	}

	// normalizando o formato do campo dhCons para incluir a timezone: (yyyy-mm-ddThh:nn:ss-03:00)
	re := regexp.MustCompile(`(<dhCons>[\d]{4}-[\d]{2}-[\d]{2}T[\d]{2}:[\d]{2}:[\d]{2})(<\/dhCons>)`)
	xmlfile2 := re.ReplaceAll(xmlfile, []byte("$1-03:00$2"))

	// normalizando o formato dos campos dIniAtiv, dUltSit e dBaixa para civil.Date (yyyy-mm-dd)
	re = regexp.MustCompile(`(<dIniAtiv>[\d]{4}-[\d]{2}-[\d]{2})[\dT\-:Z]+(<\/dIniAtiv>)`)
	xmlfile2 = re.ReplaceAll(xmlfile2, []byte("$1$2"))
	re = regexp.MustCompile(`(<dUltSit>[\d]{4}-[\d]{2}-[\d]{2})[\dT\-:Z]+(<\/dUltSit>)`)
	xmlfile2 = re.ReplaceAll(xmlfile2, []byte("$1$2"))
	re = regexp.MustCompile(`(<dBaixa>[\d]{4}-[\d]{2}-[\d]{2})[\dT\-:Z]+(<\/dBaixa>)`)
	xmlfile2 = re.ReplaceAll(xmlfile2, []byte("$1$2"))

	var ret RetConsCad
	err = xml.Unmarshal(xmlfile2, &ret)
	if err != nil {
		return RetConsCad{}, xmlfile, fmt.Errorf("Erro na desserialização do arquivo XML: %w. Arquivo: %s", err, xmlfile)
	}

	return ret, xmlfile, nil
}

// Função auxiliar para executar a ConsCad.Consulta()
func ConsultaCad(ie string, cnpj string, cpf string, cUF int, tpAmb TAmb, client *http.Client, optReq ...func(req *http.Request)) (RetConsCad, []byte, error) {
	cons := ConsCad{}
	cons.Versao = VerConsCad
	cons.InfCons.XServ = "CONS-CAD"
	cons.InfCons.IE = ie
	cons.InfCons.CNPJ = cnpj
	cons.InfCons.CPF = cpf
	cons.InfCons.UF = GetUF(cUF)

	return cons.Consulta(tpAmb, client, optReq...)
}
