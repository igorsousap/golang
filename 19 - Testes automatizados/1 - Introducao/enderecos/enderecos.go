package enderecos

import "strings"

// tipoEndereco verifica se um endereco tem um tipo valido e retorna
func tipoEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecoMinusculo := strings.ToLower(endereco)
	primeiraPalavraEndereco := strings.Split(enderecoMinusculo, " ")[0]

	enderecoTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTipoValido = true
		}
	}

	if enderecoTipoValido {
		return primeiraPalavraEndereco
	}

	return "Tipo Invalido"

}
