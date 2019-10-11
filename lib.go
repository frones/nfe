package gonfe

import (
	"fmt"
	"strconv"
	"time"
	"unicode"
)

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

func validaCNPJ(cnpj string) bool {
	return true

}

func GetChaveInfo(DFeChave string) (string, int, int, string, string, int, int, string, int, error) {
	if (len(DFeChave) != 44) || (!isNumber(DFeChave)) {
		return "", 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}

	cUF := DFeChave[0:2]
	ano, err := strconv.Atoi(DFeChave[2:4])
	if err != nil {
		return "", 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	mes, err := strconv.Atoi(DFeChave[4:6])
	if err != nil {
		return "", 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	cnpj := DFeChave[6:20]
	mod := DFeChave[20:22]
	serie, err := strconv.Atoi(DFeChave[22:25])
	if err != nil {
		return "", 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	numNF, err := strconv.Atoi(DFeChave[25:34])
	if err != nil {
		return "", 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}
	tipoEmi := DFeChave[34:35]
	cNF, err := strconv.Atoi(DFeChave[35:43])
	if err != nil {
		return "", 0, 0, "", "", 0, 0, "", 0, fmt.Errorf("Chave de Acesso inválida: %s!", DFeChave)
	}

	return cUF, ano, mes, cnpj, mod, serie, numNF, tipoEmi, cNF, nil
}

func ValidaChaveDeAcesso(DFeChave string) bool {
	if (len(DFeChave) != 44) || (!isNumber(DFeChave)) {
		return false
	}

	sum := 0
	for i, c := range DFeChave[:len(DFeChave)-1] {
		n := int(c - '0')
		sum += n * (((len(DFeChave) - 2 - i) % 8) + 2)
	}
	dv := 11 - (sum % 11)
	if dv >= 10 {
		dv = 0
	}

	if strconv.Itoa(dv) != DFeChave[len(DFeChave)-1:] {
		return false
	}

	cUF, ano, mes, cnpj, mod, _, numNF, _, _, err := GetChaveInfo(DFeChave)
	if err != nil {
		return false
	}

	if (cUF != "11") && (cUF != "12") && (cUF != "13") && (cUF != "14") && (cUF != "15") && (cUF != "16") && (cUF != "17") &&
		(cUF != "21") && (cUF != "22") && (cUF != "23") && (cUF != "24") && (cUF != "25") && (cUF != "26") && (cUF != "27") && (cUF != "28") && (cUF != "29") &&
		(cUF != "31") && (cUF != "32") && (cUF != "33") && (cUF != "35") &&
		(cUF != "41") && (cUF != "42") && (cUF != "43") &&
		(cUF != "50") && (cUF != "51") && (cUF != "52") && (cUF != "53") {
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
