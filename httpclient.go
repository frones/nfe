package gonfe

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/spacemonkeygo/openssl" //usando openssl no lugar da crypto/tls por alguma incompatibilidade com a Sefaz-RS que não consegui determinar a causa. Posteriormente pretendo investigar mais a fundo e se necessário abrir um bug report.
)

const defaultTimeout = 15 * time.Second
const defaultUserAgent = "GoNFe/0.1"

// NewHTTPClient cria um http.Client com todas as configurações necessárias para comunicação com as Sefazes.
//
// O certificado digital para usar com essa biblioteca pode ser gerado a partir de um certificado A1 da seguinte maneira:
//   openssl pkcs12 -in certificado.pfx -out ~/client.pem -clcerts -nokeys -nodes
//   openssl pkcs12 -in certificado.pfx -out ~/key.pem -nocerts -nodes
func NewHTTPClient(certFile string, certKeyFile string) (*http.Client, error) {
	pem, err := ioutil.ReadFile(certFile)
	if err != nil {
		return nil, fmt.Errorf("Erro na leitura do arquivo PEM do certificado. Detalhes: %w", err)
	}
	cert, err := openssl.LoadCertificateFromPEM(pem)
	if err != nil {
		return nil, fmt.Errorf("Erro ao carregar o certificado do arquivo PEM. Detalhes: %w", err)
	}

	pem, err = ioutil.ReadFile(certKeyFile)
	if err != nil {
		return nil, fmt.Errorf("Erro na leitura do arquivo PEM da chave. Detalhes: %w", err)
	}
	key, err := openssl.LoadPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, fmt.Errorf("Erro ao carregar a chave do arquivo PEM. Detalhes: %w", err)
	}

	ctx, err := openssl.NewCtx()
	if err != nil {
		return nil, fmt.Errorf("Erro na criação do context SSL. Detalhes: %w", err)
	}
	err = ctx.UseCertificate(cert)
	if err != nil {
		return nil, fmt.Errorf("Erro ao adicionar o certificado ao context. Detalhes: %w", err)
	}
	err = ctx.UsePrivateKey(key)
	if err != nil {
		return nil, fmt.Errorf("Erro ao adicionar a chave ao context. Detalhes: %w", err)
	}

	client := http.Client{
		Timeout: defaultTimeout,
		Transport: &http.Transport{
			DialTLS: func(network, address string) (net.Conn, error) {
				return openssl.Dial(network, address, ctx, 0)
			},
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

	xmlfile, err = getSoapEnvelope(xmlfile, xmlns)
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

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Falha na consulta à receita (%s): %d", url, resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	xmlfile, err = readSoapEnvelope(body)
	if err != nil {
		return nil, err
	}

	return xmlfile, err
}
