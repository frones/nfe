# nfe

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

	"github.com/frones/nfe"
)

func main() {
	client, err := nfe.NewHTTPClient("~/client.pem", "~/key.pem")
	if err != nil {
		fmt.Printf("Erro na criação do HTTP Client e leitura do certificado. Detalhes: %v\n", err)
		return
	}
	ret, xmlfile, err := nfe.ConsultaNFe("12345678901234567890123456789012345678901234", nfe.Homologacao, client, func(req *http.Request) {req.Header.Set("User-Agent", "MyUA/1.0")})
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

## Problemas de comunicação com a Sefaz-RS e ambientes virtuais SV-RS

Usando a `crypto/tls` padrão do Go, foi observado um problema intermitente de comunicação com os ambientes da Sefaz-RS, com resposta 403 sendo retornada. O problema acontece porque a `crypto/tls` não envia o certificado durante o handshake quando a `CertificateRequest` do servidor especifica autoridades certificadoras que não batem com a CA do certificado [[source](https://github.com/golang/go/blob/79d4defa75a26dd975c6ba3ac938e0e414dfd3e9/src/crypto/tls/common.go#L1320-L1347)]. Outras Sefazes não enviam uma lista de CAs permitidas, não apresentando esse problema. Mesmo a Sefaz-RS, em algumas requests não envia lista de CAs permitidas, fazendo com que o problema seja intermitente.

Durante depuração, identifiquei que de fato, a CA do meu certificado não estava na lista de CAs permitidas pelo servidor, de maneira que faz sentido que a `crypto/tls` não envie, mas não fica claro porque outras bibliotecas, como a OpenSSL sempre funcionam.

Um workaround possível é compilar com uma versão alterada da `crypto/tls` sempre ignorando a instrução de CAs permitidas, mas isso pode ter outras consequencias indesejadas.