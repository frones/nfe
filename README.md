# gonfe
Bibliotecas para geração, validação, assinatura e transmissão de XMLs da NFe. Pretendo desenvolver essa biblioteca conforme a necessidade surgir. Por enquanto apenas a consulta está disponível.

## Extraindo chaves de um certificado A1 (.pfx) para uso com o biblioteca
```
openssl pkcs12 -in certificado.pfx -out ~/client.key -nocerts -nodes
openssl pkcs12 -in certificado.pfx -out ~/client.pem -clcerts -nokeys -nodes
```

## Consulta NFe
### Exemplo
```go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/frones/gonfe"
)

func main() {
	var cons gonfe.ConsSitNFe
	cons.ChNFe = "1234567890123456789012345678901234"
	cons.TpAmb = gonfe.Homologacao
	ret, xmlfile, err := cons.Consulta("~/client.pem", "~/client.key")
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
