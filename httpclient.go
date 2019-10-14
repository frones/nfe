package gonfe

import (
	"bytes"
	"crypto/tls"
	"net/http"
	"time"
)

const defaultTimeout = 15 * time.Second
const defaultUserAgent = "GoNFe/0.1"

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

func newRequest(url string, body []byte) (*http.Request, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/soap+xml; charset=utf-8")
	req.Header.Set("User-Agent", defaultUserAgent)

	return req, nil
}
