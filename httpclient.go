package nfe

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const defaultTimeout = 15 * time.Second
const defaultUserAgent = "GoNFe/0.1"

// NewHTTPClient cria um http.Client com todas as configurações necessárias para comunicação com as Sefazes.
//
// O certificado digital para usar com essa biblioteca pode ser gerado a partir de um certificado A1 da seguinte maneira:
//
//	openssl pkcs12 -in certificado.pfx -out ~/client.pem -clcerts -nokeys -nodes
//	openssl pkcs12 -in certificado.pfx -out ~/key.pem -nocerts -nodes
func NewHTTPClient(certFile string, certKeyFile string) (*http.Client, error) {
	tlsConfig := tls.Config{}

	cert, err := tls.LoadX509KeyPair(certFile, certKeyFile)
	if err != nil {
		return nil, fmt.Errorf("Erro no carregamento do certificado digital. Detalhes: %w", err)
	}
	tlsConfig.Certificates = []tls.Certificate{cert}

	caCertPool, err := x509.SystemCertPool()
	if err != nil {
		return nil, fmt.Errorf("Erro no carregamento da cadeia de certificados do sistema. Detalhes: %w", err)
	}
	tlsConfig.RootCAs = caCertPool
	tlsConfig.Renegotiation = tls.RenegotiateOnceAsClient

	client := http.Client{
		Timeout: defaultTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tlsConfig,
		},
	}

	return &client, nil
}

// newRequest é uma função usada internamente para criar a requisição e já definir alguns parâmetros default. Para personalizar a sua requisição (por exemplo User-Agent) ver o parâmetro optReq da ConsSitNFe.Consulta().
func newRequest(url string, soapAction string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	if soapAction != "" {
		req.Header.Set("SOAPAction", soapAction)
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("User-Agent", defaultUserAgent)

	return req, nil
}

// sendRequest é uma função que se encarrega de fazer o envelopamento da requisição, enviar pra Sefaz com certificado digital e desenvelopar o retorno.
func sendRequest(obj interface{}, url string, xmlns string, soapAction string, client *http.Client, optReq ...func(req *http.Request)) ([]byte, error) {
	xmlfile, err := xml.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("Erro na geração do XML de requisição. Detalhes: %w", err)
	}

	if url == urlConsCadMT {
		xmlfile, err = getSoapEnvelopeConsCadMT(xmlfile, xmlns)
	} else {
		xmlfile, err = getSoapEnvelope(xmlfile, xmlns)
	}
	if err != nil {
		return nil, fmt.Errorf("Erro na geração do envelope SOAP. Detalhes: %w", err)
	}
	xmlfile = []byte(append([]byte(xml.Header), xmlfile...))

	req, err := newRequest(url, soapAction, xmlfile)
	if err != nil {
		return nil, fmt.Errorf("Erro na criação da requisição (http.Request) para a URL %s. Detalhes: %w", url, err)
	}
	for _, opt := range optReq {
		opt(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Erro na requisição ao WebService %s. Detalhes: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return nil, &WSError{url, resp.StatusCode, resp.Status, string(body)}
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Erro na leitura do corpo da resposta: %w", err)
	}

	if url == urlConsCadMT {
		xmlfile, err = readSoapEnvelopeConsCadMT(body)
	} else if url == urlConsCadMG {
		xmlfile, err = readSoapEnvelopeConsCadMG(body)
	} else {
		xmlfile, err = readSoapEnvelope(body)
	}
	if err != nil {
		return nil, err
	}

	return xmlfile, err
}
