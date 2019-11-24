package nfe

import (
	"encoding/xml"
	"fmt"
)

// Envelope representa o XML do envelope SOAP que será usado na comunicação.
type Envelope struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Body    struct {
		NfeDadosMsg struct {
			Xmlns string `xml:"xmlns,attr"`
			Value []byte `xml:",innerxml"`
		} `xml:"nfeDadosMsg"`
	} `xml:"soap12:Body"`
}

// EnvelopeResult representa o XML do envelope SOAP de retorno da requisição.
type EnvelopeResult struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		NfeResultMsg struct {
			Value []byte `xml:",innerxml"`
		} `xml:"nfeResultMsg"`
	} `xml:"Body"`
}

// getSoapEnvelope envelopa um XML fornecido de acordo com o padrão SOAP.
func getSoapEnvelope(msg []byte, xmlns string) ([]byte, error) {
	var env Envelope

	env.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.Xsd = "http://www.w3.org/2001/XMLSchema"
	env.Soap12 = "http://www.w3.org/2003/05/soap-envelope"
	env.Body.NfeDadosMsg.Xmlns = xmlns
	env.Body.NfeDadosMsg.Value = msg

	return xml.Marshal(env)
}

// readSoapEnvelope extrai o XML de resposta de um envelope SOAP de retorno.
func readSoapEnvelope(msg []byte) ([]byte, error) {
	var env EnvelopeResult
	err := xml.Unmarshal(msg, &env)
	if err != nil {
		return nil, fmt.Errorf("Erro na desserialização do arquivo XML: %w. Arquivo: %s", err, msg)
	}
	return env.Body.NfeResultMsg.Value, nil
}
