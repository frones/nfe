package nfe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Esse exemplo mostra todos os passos para se fazer uma consulta de protocolo na Sefaz. Desde a criação de um novo http.Client (através da NewHTTPClient) até a personalização do User-Agent por meio do parâmetro optReq.
func ExampleConsultaNFe() {
	client, err := nfe.NewHTTPClient("~/client.pem", "~/key.pem")
	if err != nil {
		fmt.Printf("Erro na criação do HTTP Client e leitura do certificado. Detalhes: %v\n", err)
		return
	}
	ret, xmlfile, err := nfe.ConsultaNFe("12345678901234567890123456789012345678901234", nfe.Homologacao, client, func(req *http.Request) { req.Header.Set("User-Agent", "MyUA/1.0") })
	if err != nil {
		fmt.Printf("Erro na consulta da chave de acesso. Detalhes: %v\n", err)
		return
	}

	fmt.Printf("Objeto RetConsSitNFe: %v\n", ret)
	fmt.Printf("Arquivo XML: %v\n", string(xmlfile))

	jsonfile, err := json.Marshal(ret)
	if err != nil {
		fmt.Printf("Erro na serialização do json. Detalhes: %v\n", err)
		return
	}
	fmt.Printf("Arquivo JSON: %v\n", string(jsonfile))
}

func ExampleConsSitNFe_Consulta() {
	client, err := nfe.NewHTTPClient("~/client.pem", "~/key.pem")
	if err != nil {
		fmt.Printf("Erro na criação do HTTP Client e leitura do certificado. Detalhes: %v\n", err)
		return
	}

	cons := nfe.ConsSitNFe{
		Versao: nfe.VerConsSitNFe,
		TpAmb:  nfe.Homologacao,
		XServ:  "CONSULTAR",
		ChNFe:  "12345678901234567890123456789012345678901234",
	}

	ret, xmlfile, err := cons.Consulta(client)
	if err != nil {
		fmt.Printf("Erro na consulta da chave de acesso. Detalhes: %v\n", err)
		return
	}

	fmt.Printf("%v\n\n", ret)
	fmt.Printf("%s\n", xmlfile)
}

// Esse exemplo mostra todos os passos para se fazer uma consulta de status do serviço na Sefaz. Desde a criação de um novo http.Client (através da NewHTTPClient) até a personalização do User-Agent por meio do parâmetro optReq.
func ExampleConsultaStatServ() {
	client, err := nfe.NewHTTPClient("~/client.pem", "~/key.pem")
	if err != nil {
		fmt.Printf("Erro na criação do HTTP Client e leitura do certificado. Detalhes: %v\n", err)
		return
	}
	ret, xmlfile, err := nfe.ConsultaStatServ(35, nfe.Homologacao, client, func(req *http.Request) { req.Header.Set("User-Agent", "MyUA/1.0") })
	if err != nil {
		fmt.Printf("Erro na consulta da chave de acesso. Detalhes: %v\n", err)
		return
	}

	fmt.Printf("Objeto RetConsStatServ: %v\n", ret)
	fmt.Printf("Arquivo XML: %v\n", string(xmlfile))

	jsonfile, err := json.Marshal(ret)
	if err != nil {
		fmt.Printf("Erro na serialização do json. Detalhes: %v\n", err)
		return
	}
	fmt.Printf("Arquivo JSON: %v\n", string(jsonfile))
}

func ExampleConsStatServ_Consulta() {
	client, err := nfe.NewHTTPClient("~/client.pem", "~/key.pem")
	if err != nil {
		fmt.Printf("Erro na criação do HTTP Client e leitura do certificado. Detalhes: %v\n", err)
		return
	}

	cons := nfe.ConsStatServ{
		Versao: nfe.VerConsStatServ,
		TpAmb:  nfe.Homologacao,
		XServ:  "STATUS",
		CUF:    35,
	}

	ret, xmlfile, err := cons.Consulta(client)
	if err != nil {
		fmt.Printf("Erro na consulta da chave de acesso. Detalhes: %v\n", err)
		return
	}

	fmt.Printf("%v\n\n", ret)
	fmt.Printf("%s\n", xmlfile)
}
