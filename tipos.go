package gonfe

import (
	"time"
)

type TAmb int

const (
	Producao    TAmb = 1
	Homologacao TAmb = 2
)

type TWebService int

const (
	ConsultaStatus TWebService = iota
	ConsultaProtocolo
	ConsultaCadastro
	Autorizacao
	RetAutorizacao
	Evento
	Inutilizacao
)

type ProtNFe struct {
	Versao  string `json:"-" xml:"versao,attr"`
	InfProt struct {
		TpAmb    TAmb      `json:"tpAmb" xml:"tpAmb"`
		VerAplic string    `json:"verAplic" xml:"verAplic"`
		ChNFe    string    `json:"chNFe" xml:"chNFe"`
		DhRecbto time.Time `json:"dhRecbto" xml:"dhRecbto"`
		NProt    string    `json:"nProt" xml:"nProt"`
		DigVal   string    `json:"digVal" xml:"digVal"`
		CStat    int       `json:"cStat" xml:"cStat"`
		XMotivo  string    `json:"xMotivo" xml:"xMotivo"`
	} `json:"infProt" xml:"infProt"`
}
