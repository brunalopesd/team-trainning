package atbash

import (
	"regexp"
	"strings"
)

/*O texto cifrado é escrito em grupos de comprimento fixo,
sendo o tamanho tradicional do grupo de 5 letras, deixando os números inalterados e a pontuação é excluída
*/

func Atbash(text string) string {
	normalizedText := strings.ToLower(text)
	reg := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	normalizedText = reg.ReplaceAllString(normalizedText, "")

	result := []rune{}

	for i, char := range normalizedText {
		if i > 1 && i%5 == 0 {
			result = append(result, ' ')
		}
		if char >= 'a' && char <= 'z' {
			cypherAlphabet := char - 'a' //ajusta o indice, subtraindo o valor ascii do charctere 'a'
			// (25 - cypherAlphabet) inverte a posição do alfabeto se char é 'a' (com indice 0) a = 0 -> 25 - 0 = 25 (última posição, que corresponde a 'z').
			result = append(result, (25 - cypherAlphabet))
		} else {
			result = append(result, char)
		}
	}
	return string(result)

}
