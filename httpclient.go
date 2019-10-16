package gonfe

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"time"
)

const defaultTimeout = 15 * time.Second
const defaultUserAgent = "GoNFe/0.1"

// NewHTTPClient cria um http.Client com todas as configurações necessárias para comunicação com as Sefazes.
//
// O certificado digital para usar com essa biblioteca pode ser gerado a partir de um certificado A1 da seguinte maneira:
//   openssl pkcs12 -in certificado.pfx -out ~/client.pem -clcerts -nokeys -nodes
//   openssl pkcs12 -in certificado.pfx -out ~/key.pem -nocerts -nodes
func NewHTTPClient(certFile string, certKeyFile string) (*http.Client, error) {
	cert, err := tls.LoadX509KeyPair(certFile, certKeyFile)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: defaultTimeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates:       []tls.Certificate{cert},
				InsecureSkipVerify: true,
				Renegotiation:      tls.RenegotiateFreelyAsClient,
			},
		},
	}

	return &client, nil
}

// newRequest é uma função usada internamente para criar a requisição e já definir alguns parâmetros default. Para personalizar a sua requisição (por exemplo User-Agent) ver o parâmetro optReq da ConsSitNFe.Consulta().
func newRequest(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("User-Agent", defaultUserAgent)

	return req, nil
}
