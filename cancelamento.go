package nfe

import (
	"time"
)

// RetCancNFe representa o XML de retorno da Sefaz do cancelamento da NFe. Não é mais usado, tendo sido substituído pelos eventos (EventoNFe), mas ainda pode ser retornado em uma consulta de protocolo (ConsSitNFe) de notas antigas.
type RetCancNFe struct {
	Versao  string `json:"-" xml:"versao,attr"`
	InfCanc struct {
		TpAmb    TAmb      `json:"tpAmb" xml:"tpAmb"`
		VerAplic string    `json:"verAplic" xml:"verAplic"`
		CStat    int       `json:"cStat" xml:"cStat"`
		XMotivo  string    `json:"xMotivo" xml:"xMotivo"`
		CUF      int       `json:"cUF" xml:"cUF"`
		ChNFe    string    `json:"chNFe" xml:"chNFe"`
		DhRecbto time.Time `json:"dhRecbto" xml:"dhRecbto"`
		NProt    string    `json:"nProt" xml:"nProt"`
	} `json:"infCanc" xml:"infCanc"`
}
