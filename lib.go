//O package gonfe fornece funções para fazer toda a comunicação com as Sefazes no âmbito da NFe.
package gonfe

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// isNumber verifica se todos os caracteres da string são dígitos numéricos (usado internamente para validação da chave de acesso).
func isNumber(s string) bool {
	if s == "" {
		return false
	}

	for _, d := range s {
		if !unicode.IsDigit(d) {
			return false
		}
	}

	return true
}

// validaCNPJ verifica se o CNPJ informado é válido.
func validaCNPJ(cnpj string) bool {
	cnpj = strings.Replace(cnpj, ".", "", -1)
	cnpj = strings.Replace(cnpj, "-", "", -1)
	cnpj = strings.Replace(cnpj, "/", "", -1)

	if (len(cnpj) != 14) || (!isNumber(cnpj)) {
		return false
	}

	sum := 0
	for i, c := range cnpj[:len(cnpj)-2] {
		sum += int(c-'0') * (((len(cnpj) - 3 - i) % 8) + 2)
	}
	dv := ((sum * 10) % 11) % 10
	if dv != int(cnpj[12]-'0') {
		return false
	}

	sum = 0
	for i, c := range cnpj[:len(cnpj)-1] {
		sum += int(c-'0') * (((len(cnpj) - 2 - i) % 8) + 2)
	}
	dv = ((sum * 10) % 11) % 10
	if dv != int(cnpj[13]-'0') {
		return false
	}

	return true
}

// GetChaveInfo extrai todas as informações que estão embutidas em uma chave de acesso da NFe:
//
// cUF, Ano, Mes, CNPJ, Modelo (55/65), Número da NFe, tpEmis e cNF.
func GetChaveInfo(DFeChave string) (int, int, int, string, string, int, int, string, int, error) {
	if (len(DFeChave) != 44) || (!isNumber(DFeChave)) {
		return 0, 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}

	cUF, err := strconv.Atoi(DFeChave[0:2])
	if err != nil {
		return 0, 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	ano, err := strconv.Atoi(DFeChave[2:4])
	if err != nil {
		return 0, 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	mes, err := strconv.Atoi(DFeChave[4:6])
	if err != nil {
		return 0, 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	cnpj := DFeChave[6:20]
	mod := DFeChave[20:22]
	serie, err := strconv.Atoi(DFeChave[22:25])
	if err != nil {
		return 0, 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	numNF, err := strconv.Atoi(DFeChave[25:34])
	if err != nil {
		return 0, 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	tipoEmi := DFeChave[34:35]
	cNF, err := strconv.Atoi(DFeChave[35:43])
	if err != nil {
		return 0, 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}

	return cUF, ano, mes, cnpj, mod, serie, numNF, tipoEmi, cNF, nil
}

// ValidaChaveDeAcesso verifica se a chave de acesso fornecida é válida, através dos seguintes critérios:
//   * Tamanho = 44 e conteúdo numérico
//   * Dígito verificador consistente
//   * cUF corresponde a um item da tabela do IBGE
//   * Mes/Ano válidos, posteriores a 01/2006 e não posteriores ao ano atual
//   * CNPJ válido e diferente de 00000000000000
//   * Modelo igual a 55 ou 65 ou 67
//   * Número da NF diferente de zero
func ValidaChaveDeAcesso(DFeChave string) bool {
	if (len(DFeChave) != 44) || (!isNumber(DFeChave)) {
		return false
	}

	sum := 0
	for i, c := range DFeChave[:len(DFeChave)-1] {
		n := int(c - '0')
		sum += n * (((len(DFeChave) - 2 - i) % 8) + 2)
	}
	dv := ((sum * 10) % 11) % 10

	if strconv.Itoa(dv) != DFeChave[len(DFeChave)-1:] {
		return false
	}

	cUF, ano, mes, cnpj, mod, _, numNF, _, _, err := GetChaveInfo(DFeChave)
	if err != nil {
		return false
	}

	if (cUF != 11) && (cUF != 12) && (cUF != 13) && (cUF != 14) && (cUF != 15) && (cUF != 16) && (cUF != 17) &&
		(cUF != 21) && (cUF != 22) && (cUF != 23) && (cUF != 24) && (cUF != 25) && (cUF != 26) && (cUF != 27) && (cUF != 28) && (cUF != 29) &&
		(cUF != 31) && (cUF != 32) && (cUF != 33) && (cUF != 35) &&
		(cUF != 41) && (cUF != 42) && (cUF != 43) &&
		(cUF != 50) && (cUF != 51) && (cUF != 52) && (cUF != 53) {
		return false
	}

	if (ano < 6) || (2000+ano > time.Now().Year()) {
		return false
	}

	if (mes < 1) || (mes > 12) {
		return false
	}

	if (cnpj == "00000000000000") || (!validaCNPJ(cnpj)) {
		return false
	}

	if (mod != "55") && (mod != "65") && (mod != "67") {
		return false
	}

	if numNF == 0 {
		return false
	}

	return true
}
