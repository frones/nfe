# gonfe
Bibliotecas para geração, validação, assinatura e transmissão de XMLs da NFe. Pretendo desenvolver essa biblioteca conforme a necessidade surgir. Por enquanto apenas as consultas de Status, NFe e Cadastro estão disponíveis.

## Extraindo chaves de um certificado A1 (.pfx) para uso com o biblioteca
```
openssl pkcs12 -in certificado.pfx -out ~/client.pem -clcerts -nokeys -nodes
openssl pkcs12 -in certificado.pfx -out ~/key.pem -nocerts -nodes
```

## Consulta NFe
### Exemplo
```go
package main

import (
	"encoding/json"
	"net/http"
	"fmt"

	"github.com/frones/gonfe"
)

func main() {
	client, err := gonfe.NewHTTPClient("~/client.pem", "~/key.pem")
	if err != nil {
		fmt.Printf("Erro na criação do HTTP Client e leitura do certificado. Detalhes: %v\n", err)
		return
	}
	ret, xmlfile, err := gonfe.ConsultaNFe("12345678901234567890123456789012345678901234", gonfe.Homologacao, client, func(req *http.Request) {req.Header.Set("User-Agent", "MyUA/1.0")})
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
```
