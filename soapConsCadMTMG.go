package nfe

import (
	"encoding/xml"
	"fmt"
)

// EnvelopeConsCadMT representa o XML do envelope SOAP que será usado na comunicação com a consulta de cadastro do MT.
type EnvelopeConsCadMT struct {
	XMLName xml.Name `xml:"soap12:Envelope"`
	Xsi     string   `xml:"xmlns:xsi,attr"`
	Xsd     string   `xml:"xmlns:xsd,attr"`
	Soap12  string   `xml:"xmlns:soap12,attr"`
	Body    struct {
		ConsultaCadastro struct {
			Xmlns       string `xml:"xmlns,attr"`
			NfeDadosMsg struct {
				Value []byte `xml:",innerxml"`
			} `xml:"nfeDadosMsg"`
		} `xml:"consultaCadastro"`
	} `xml:"soap12:Body"`
}

// EnvelopeResultConsCadMG representa o XML do envelope SOAP de retorno da requisição ConsCad em MG, que é diferente de todas as outras respostas do sistema NFe.
type EnvelopeResultConsCadMG struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		ConsultaCadastro4Result struct {
			Value []byte `xml:",innerxml"`
		} `xml:"consultaCadastro4Result"`
	} `xml:"Body"`
}

// EnvelopeResultConsCadMT representa o XML do envelope SOAP de retorno da requisição ConsCad em MT, que é diferente de todas as outras respostas do sistema NFe.
type EnvelopeResultConsCadMT struct {
	XMLName xml.Name `xml:"Envelope"`
	Body    struct {
		NfeResultMsg struct {
			ConsultaCadastroResult struct {
				Value []byte `xml:",innerxml"`
			} `xml:"consultaCadastroResult"`
		} `xml:"nfeResultMsg"`
	} `xml:"Body"`
}

// getSoapEnvelopeConsCadMT envelopa um XML fornecido de acordo com o padrão SOAP exclusivo de MT.
func getSoapEnvelopeConsCadMT(msg []byte, xmlns string) ([]byte, error) {
	var env EnvelopeConsCadMT

	env.Xsi = "http://www.w3.org/2001/XMLSchema-instance"
	env.Xsd = "http://www.w3.org/2001/XMLSchema"
	env.Soap12 = "http://www.w3.org/2003/05/soap-envelope"
	env.Body.ConsultaCadastro.Xmlns = xmlns
	env.Body.ConsultaCadastro.NfeDadosMsg.Value = msg

	return xml.Marshal(env)
}

// readSoapEnvelopeConsCadMT extrai o XML de resposta de um envelope SOAP de retorno da consulta de cadastro de MT.
func readSoapEnvelopeConsCadMT(msg []byte) ([]byte, error) {
	var env EnvelopeResultConsCadMT
	err := xml.Unmarshal(msg, &env)
	if err != nil {
		return nil, fmt.Errorf("Erro na desserialização do arquivo XML: %w. Arquivo: %s", err, msg)
	}
	return env.Body.NfeResultMsg.ConsultaCadastroResult.Value, nil
}

// readSoapEnvelopeConsCadMG extrai o XML de resposta de um envelope SOAP de retorno da consulta de cadastro de MG.
func readSoapEnvelopeConsCadMG(msg []byte) ([]byte, error) {
	var env EnvelopeResultConsCadMG
	err := xml.Unmarshal(msg, &env)
	if err != nil {
		return nil, fmt.Errorf("Erro na desserialização do arquivo XML: %w. Arquivo: %s", err, msg)
	}
	return env.Body.ConsultaCadastro4Result.Value, nil
}
