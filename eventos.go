package nfe

import (
	"encoding/xml"
	"time"

	"github.com/amdonov/xmlsig"
)

// ProcEventoNFe representa o XML que contem tanto a requisição (EventoNFe) quanto o retorno da Sefaz (RetEventoNFe), e poderá vir dentro de consultas de status (ConsSitNFe).
type ProcEventoNFe struct {
	XMLName   xml.Name      `json:"-" xml:"procEventoNFe"`
	Versao    string        `json:"versao" xml:"versao,attr"`
	Evento    *EventoNFe    `json:"evento" xml:"http://www.portalfiscal.inf.br/nfe evento"`
	RetEvento *RetEventoNFe `json:"retEvento" xml:"retEvento"`
}

// EventoNFe representa o XML de registro de um evento junto à Sefaz.
type EventoNFe struct {
	Versao    string `json:"-" xml:"versao,attr"`
	InfEvento struct {
		Versao     string    `json:"versao,omitempty" xml:"versao,attr,omitempty"`
		ID         string    `json:"Id" xml:"Id,attr"`
		COrgao     int       `json:"cOrgao" xml:"cOrgao"`
		TpAmb      TAmb      `json:"tpAmb" xml:"tpAmb"`
		CNPJ       string    `json:"CNPJ,omitempty" xml:"CNPJ,omitempty"`
		CPF        string    `json:"CPF,omitempty" xml:"CPF,omitempty"`
		ChNFe      string    `json:"chNFe" xml:"chNFe"`
		DhEvento   time.Time `json:"dhEvento" xml:"dhEvento"`
		TpEvento   string    `json:"tpEvento" xml:"tpEvento"`
		NSeqEvento int       `json:"nSeqEvento" xml:"nSeqEvento"`
		VerEvento  string    `json:"verEvento" xml:"verEvento"`
		DetEvento  struct {
			Versao string `xml:"versao,attr"`
			Value  []byte `xml:",innerxml"`
		} `json:"-" xml:"detEvento"`
	} `json:"infEvento" xml:"infEvento"`
	Signature *xmlsig.Signature `json:"-" xml:"Signature"`
}

// RetEventoNFe representa o XML de retorno da Sefaz à solicitação de registro de evento. Normalmente será utilizado encapsulado em um ProcEventoNFe.
type RetEventoNFe struct {
	Versao    string `json:"versao" xml:"versao,attr"`
	InfEvento struct {
		ID          string    `json:"Id" xml:"Id,attr,omitempty"`
		TpAmb       TAmb      `json:"tpAmb" xml:"tpAmb"`
		VerAplic    string    `json:"verAplic" xml:"verAplic"`
		COrgao      int       `json:"cOrgao" xml:"cOrgao"`
		CStat       int       `json:"cStat" xml:"cStat"`
		XMotivo     string    `json:"xMotivo" xml:"xMotivo"`
		ChNFe       string    `json:"chNFe" xml:"chNFe"`
		TpEvento    string    `json:"tpEvento" xml:"tpEvento"`
		XEvento     string    `json:"xEvento" xml:"xEvento"`
		NSeqEvento  int       `json:"nSeqEvento" xml:"nSeqEvento"`
		CNPJDest    string    `json:"CNPJDest,omitempty" xml:"CNPJDest,omitempty"`
		CPFDest     string    `json:"CPFDest,omitempty" xml:"CPFDest,omitempty"`
		EmailDest   string    `json:"emailDest,omitempty" xml:"emailDest,omitempty"`
		DhRegEvento time.Time `json:"dhRegEvento" xml:"dhRegEvento"`
		NProt       string    `json:"nProt" xml:"nProt"`
	} `json:"infEvento" xml:"infEvento"`
}
